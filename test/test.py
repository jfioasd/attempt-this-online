from asyncio import gather, get_event_loop, sleep
from contextlib import asynccontextmanager
from os import environ
import subprocess
from time import monotonic
from websockets import connect, ConnectionClosed
from msgpack import loads, dumps
from pytest import mark, raises
from pytest_asyncio import fixture

KiB = 1024
sec = 1e9  # ns
SIGKILL = 9
UNSUPPORTED_DATA = 1003
TOO_LARGE = 1009
POLICY_VIOLATION = 1008

url = environ["URL"]
slow = mark.skipif(bool(environ.get("FAST")), reason="takes too long and FAST option set")


# use only for simple tests which only need one connection and don't handle connection exceptions
@fixture
async def c():
    async with connect(url) as conn:
        yield conn


def req(code, input="", options=(), arguments=(), language="zsh", timeout=60, hook=None):
    d = {
        "language": language,
        "code": code,
        "input": input,
        "arguments": arguments,
        "options": options,
        "timeout": timeout,
    }
    if hook:
        hook(d)
    return dumps(d)


async def test_noop():
    async with connect(url):
        pass


async def test_basic_execution(c):
    await c.send(req(""))
    r = loads(await c.recv())
    assert r.keys() == {"Done"}
    r = r["Done"]
    # check stats values are reasonable
    assert r["user"] + r["kernel"] <= r["real"]
    assert r.pop("real") < 0.1 * sec
    assert r.pop("kernel") < 0.1 * sec
    assert r.pop("user") < 0.1 * sec
    assert 1000 < r.pop("max_mem") < 100_000  # KiB
    assert 0 <= r.pop("waits") < 1000
    assert 0 <= r.pop("preemptions") < 1000
    assert 0 <= r.pop("major_page_faults") < 1000
    assert 0 <= r.pop("minor_page_faults") < 1000
    assert 0 <= r.pop("input_ops") < 100000
    assert 0 <= r.pop("output_ops") < 100
    assert r == {
        "timed_out": False,
        "status_type": "exited",
        "status_value": 0,
    }


async def test_connection_reuse(c):
    await test_basic_execution(c)
    await test_basic_execution(c)


@slow
async def test_code(c):
    start = monotonic()
    await c.send(req("sleep 1"))
    assert loads(await c.recv()).keys() == {"Done"}
    assert 1 < monotonic() - start < 1.1


@slow
async def test_parallelism():
    async def inner():
        async with connect(url) as c:
            await test_code(c)

    start = monotonic()
    await gather(inner(), inner())
    assert 1 < monotonic() - start < 1.2


async def test_stdout(c):
    await c.send(req("echo hello"))
    assert loads(await c.recv()) == {"Stdout": b"hello\n"}
    assert loads(await c.recv()).keys() == {"Done"}


async def test_stderr(c):
    await c.send(req("echo hello >&2"))
    assert loads(await c.recv()) == {"Stderr": b"hello\n"}
    assert loads(await c.recv()).keys() == {"Done"}


async def test_stdin(c):
    await c.send(req("rev", input="hello"))
    assert loads(await c.recv()) == {"Stdout": b"olleh"}
    assert loads(await c.recv()).keys() == {"Done"}


async def test_args(c):
    await c.send(req("echo $@", arguments=["foo", "bar"]))
    assert loads(await c.recv()) == {"Stdout": b"foo bar\n"}
    assert loads(await c.recv()).keys() == {"Done"}


async def test_options(c):
    await c.send(req("echo $-", options=["-F"]))
    assert b"F" in loads(await c.recv())["Stdout"]
    assert loads(await c.recv()).keys() == {"Done"}


async def test_exit(c):
    await c.send(req("exit 7"))
    r = loads(await c.recv())["Done"]
    assert r["status_type"] == "exited"
    assert r["status_value"] == 7


@slow
async def test_timeout(c):
    start = monotonic()
    await c.send(req("sleep 3", timeout=1))
    r = loads(await c.recv())["Done"]
    assert 1 < monotonic() - start < 1.1
    assert r["timed_out"]
    assert r["status_type"] == "killed"
    assert r["status_value"] == SIGKILL


@asynccontextmanager
async def _test_error(msg, code=POLICY_VIOLATION, max_time=0.1):
    start = monotonic()
    with raises(ConnectionClosed) as e:
        async with connect(url) as c:
            yield c
            await c.recv()
    assert monotonic() - start < max_time
    assert e.value.code == code
    assert e.value.reason == msg


@mark.parametrize("kwargs,msg", (
    ({"timeout": 61}, "invalid request: timeout not in range 1-60: 61"),
    ({"timeout": 0}, "invalid request: timeout not in range 1-60: 0"),
    ({"timeout": -4}, "invalid request: timeout not in range 1-60: -4"),
    ({"language": "doesntexist"}, "invalid request: no such language: doesntexist"),
    ({"language": "ZSH"}, "invalid request: no such language: ZSH"),
    ({"arguments": ["null\0byte"]}, "invalid request: argument contains null byte"),
    ({"options": ["null\0byte"]}, "invalid request: argument contains null byte"),
))
async def test_invalid_request_values(kwargs, msg):
    async with _test_error(msg) as c:
        await c.send(req("sleep 1", **kwargs))


class StartsWith(str):
    def __eq__(self, other):
        if isinstance(other, str):
            return other.startswith(self)
        else:
            return NotImplemented


@mark.parametrize("kwargs", (
    {"timeout": "60"},
    {"timeout": 60.0},
    {"timeout": None},
    {"timeout": [60]},
))
async def test_invalid_request_types(kwargs):
    async with _test_error(StartsWith("invalid request:")) as c:
        await c.send(req("sleep 1", **kwargs))


async def test_invalid_request_syntax():
    async with _test_error(StartsWith("invalid request:")) as c:
        await c.send(b"not a valid msgpack message!")


async def test_incomplete_request():
    payload = req("echo hello")
    async with _test_error(StartsWith("invalid request:")) as c:
        await c.send(payload[:20])


async def test_split_request():
    payload = req("echo hello")

    async with _test_error(StartsWith("invalid request:"), max_time=1.1) as c:
        await c.send(payload[:20])
        await sleep(1)
        await c.send(payload[20:])

    async with _test_error(StartsWith("invalid request:")) as c:
        await c.send(payload[:20])
        await c.send(payload[20:])


async def test_invalid_request_data_type():
    async with _test_error("expected a binary message", UNSUPPORTED_DATA) as c:
        await c.send("not a binary message!")


@mark.xfail(True, reason="#97 is not yet fixed")
async def test_extra_junk_after_request():
    async with _test_error("invalid request: found extra data") as c:
        await c.send(req("") + b"extra junk")


async def test_too_large_request():
    s = 64 * KiB

    async with _test_error(StartsWith("invalid request:")) as c:
        await c.send(bytes(s))

    async with _test_error(f"received message of size {s + 1}, greater than size limit {s}", TOO_LARGE) as c:
        await c.send(bytes(s + 1))


@mark.parametrize("kwargs", (
    {"input": "unicode"},
    {"input": b"bytes"},
    {"input": list(b"bytes")},
    {"hook": lambda d: d.pop("timeout")},
))
async def test_valid_request_types(c, kwargs):
    await c.send(req("", **kwargs))
    assert "Done" in loads(await c.recv())


@slow
async def test_streaming(c):
    then = monotonic()
    await c.send(req("repeat 3 sleep 1 && echo hi "))
    for _ in range(3):
        assert loads(await c.recv())["Stdout"] == b"hi\n"
        now = monotonic()
        assert 0.9 < now - then < 1.1
        then = now
    assert "Done" in loads(await c.recv())
    assert monotonic() - then < 0.1


@slow
async def test_kill(c):
    await c.send(req("sleep 3"))
    await sleep(1)
    start = monotonic()
    await c.send(dumps("Kill"))
    r = loads(await c.recv())["Done"]
    assert monotonic() - start < 0.1
    assert r["status_type"] == "killed"
    assert r["status_value"] == SIGKILL


async def pgrep(*args):
    def inner():
        proc = subprocess.run(["pgrep", *args])
        match proc.returncode:
            case 0:
                return True
            case 1:
                return False
            case _:
                raise RuntimeError("pgrep failed", proc.stderr)
    return await get_event_loop().run_in_executor(None, inner)


@slow
async def test_client_close():
    start = monotonic()
    async with connect(url) as c:
        await c.send(req("sleep 5"))
        await sleep(1)
        assert await pgrep("sleep")
    await sleep(1)
    assert monotonic() - start < 2.1
    assert not await pgrep("sleep")


@slow
async def test_client_close_yes():
    async with connect(url) as c:
        await c.send(req("yes"))
    await sleep(1)
    assert not await pgrep("yes")


async def test_large_output(c):
    await c.send(req(f"dd if=/dev/zero of=/dev/stdout bs={3840 * 3} count=1 2>&-"))
    for _ in range(3):
        assert loads(await c.recv()) == {"Stdout": bytes(3840)}
    assert "Done" in loads(await c.recv())
