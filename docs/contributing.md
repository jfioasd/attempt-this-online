# Contributing
## Adding Languages
If you're not familiar with the typical GitHub (fork-edit-PR) workflow, please read the [GitHub
guide](https://guides.github.com/introduction/flow/) on the matter.

1. Create a [Docker image](https://hub.docker.com) with the toolchain for the language, and submit it to the
   [languages](https://github.com/attempt-this-online/languages) repository, based on our common base image
   `attemptthisonline/base` (which uses Arch Linux)
   - in rare cases (such as for non-open-source languages), an existing Docker image that isn't based on
     `attemptthisonline/base` may be used
   - for languages that require other languages, base the new language on an `attemptthisonline/` image for the required
     language. For example, Jelly's interpreter is written in Python, so Jelly's image is based on
     `attemptthisonline/python`
   - the build process of the Docker image should follow the general pattern of:
     - declare a Docker build argument corresponding to the language's version; this can be used as an evironment
       variable throughout the Dockerfile thereafter. This allows fast updating of the Docker images
     - install any dependencies necessary to compile the language from source
     - download and extract the source code of the language
     - compile the language's source code
     - install the compiled language (if just a few binaries are provided, they should go in `/usr/local/bin/`; if the
       installation is more complex, put it into a directory in `/opt/`)
     - clean up any downloaded or cached files that are no longer needed
   - a fairly easy-to-follow example of this is [Zsh](https://github.com/attempt-this-online/languages/blob/main/languages/zsh/Dockerfile)
   - if the language is particularly complex (or just slow) to build from source, a pre-built version can be used
     (example: [Java](https://github.com/attempt-this-online/languages/blob/main/languages/java/Dockerfile))
   - make a pull request to add it to the repository (from where it will be built and pushed to Docker Hub automatically)
2. Add the language's metadata to `languages.json`; set the *key* to an identifier-safe name for the
   language (avoid special characters); in the value, set these fields:
   - name (should be human-readable - this is what will be presented to the user in the UI)
   - image (name of the Docker image used to run the code)
   - version (set this to whatever boundary we guarantee to the user, not just the version it currently uses)
   - URL (homepage of the language)
   - SBCS (set to `true` if the language's code uses a single-byte character set; this will change the behaviour of the
     byte counter in the frontend to assume all characters comprise one byte (rather than using UTF-8))
   - se_class (provide this only if StackExchange has built-in syntax highlighting for the language; this will be added
     when a CGCC post template is generated. See [here](https://meta.stackexchange.com/q/184108) for details)
3. Create a runner script in `runners/`. Here is an example showing the general idea:

```sh
#!/bin/sh
# There are no minimal requirements for the Docker image, as long as it doesn't contain a /ATO directory. If the Docker
# image you're using doesn't have a POSIX shell, it is always available as `/ATO/bash`. If you need to use it, make sure
# to change the `#!` (shebang) line above to match that.

# Do whatever is necessary to compile and run the code.
# - code is saved in /ATO/code
# - input provided by the user is saved in /ATO/input
# - options to pass to the compiler or interpreter are stored in /ATO/options, null-terminated
# - options to pass to the program itself are stored in /ATO/arguments

# Use /ATO/yargs to substitute in the command-line options: the first argument is the replacement string, the second is
# the intput file for the arguments, and after is the program and its arguments. The replacement string indicates the
# position of the substitution.
/ATO/yargs % /ATO/options gcc % /ATO/code -o /ATO/compiled

# Note that, while the script will always start in /ATO/, you should always use absolute paths.

# The code itself should always be run in the working directory /ATO/context
cd /ATO/context

# Pass arguments to the compiled file. Also, make sure you give the program input from /ATO/input.
/ATO/yargs % /ATO/arguments /ATO/compiled % < /ATO/input

# Make sure you retain the status code of the program! If you need to do any cleanup for whatever reason, make sure to
# store a copy of the exit code and use it again.
stored_status="$?"
do_cleanup
exit "$stored_status"
```

Here is another example runner for an interpreted language instead:

```sh
#!/bin/sh

cd /ATO/context

# Use two levels of yargs to substitute in multiple sets of arguments:
/ATO/yargs %1 /ATO/options /ATO/yargs %2 /ATO/arguments python %1 /ATO/code %2 < /ATO/input
```
  - Make sure you've made the runner script executable (`chmod +x runners/path`)
  - Test your runner! It's unhelpful if you submit a broken runner
  - Make a [Pull Request](https://github.com/attempt-this-online/attempt-this-online/pulls) to add the runner for

## Making Releases
- Update version numbers in `frontend/package.json`, `Cargo.toml`, and `setup/setup`
- Upgrade dependencies (`cd frontend; npm update; cd ..; cargo update`)
- Stage and commit changes
- Tag version in git, e.g. `v0.1.2`
- Push `main` **and the new tag** to GitHub
- Build package `./build`
- Upload `setup/setup` and `dist/attempt_this_online.tar.gz` to GitHub release
- Set description etc. on GitHub release

## Backend developer instructions
The backend is written in Rust. You'll need the nightly Rust compiler and cargo.

- `src/main.rs` is the main entrypoint to the service and contains the websocket server handling code
- `src/sandbox.rs` contains the core sandbox and execution wrapper

See [Architecture](./architecture.md) for more details on how the overall system works.

This sandbox uses features specific to the Linux kernel, so you'll need to be developing on Linux.
You'll need to make sure unprivileged namespaces are enabled; instructions will vary by distribution.

### Environment
To set up a minimal development environment:

```bash
sudo mkdir -p /usr/local/{lib,share}/ATO
sudo chown $USER:$USER /usr/local/{lib,share}/ATO
mkdir /usr/local/lib/ATO/{rootfs/attemptthisonline+zsh/{proc,sys,dev,ATO},env} /usr/local/share/ATO/runners
sudo docker run --rm -it attemptthisonline/zsh \
    tar --exclude /sys --exclude /proc --exclude /dev -c / \
    | tar -xC /usr/local/lib/ATO/rootfs/attemptthisonline+zsh
printf 'PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin\0LANG=C.UTF-8\0' > /usr/local/lib/ATO/env/attemptthisonline+zsh
ln -s "$(pwd)/runners/zsh" /usr/local/share/ATO/runners/zsh
ln -s "$(which bash)" /usr/local/lib/ATO/bash
ln -s "$(pwd)/dist/attempt_this_online/yargs" /usr/local/lib/ATO/yargs
mkdir -p dist/attempt_this_online
gcc -Wall -Werror -static yargs.c -o dist/attempt_this_online/yargs
cargo build --all-targets
```

Set up the cgroup v2 pseudofilesystem. If systemd manages cgroup v2:

```bash
export ATO_CGROUP_PATH=/sys/fs/cgroup/user.slice/user-$(id -u).slice/user@$(id -u).service/ATO
mkdir -p $ATO_CGROUP_PATH
echo +memory > $ATO_CGROUP_PATH/cgroup.subtree_control
```

If your distribution doesn't use systemd with cgroup v2 (the "unified cgroup hierarchy"), then you'll have to set it up
manually.

Finally create a temporary directory owned by your user:

```
sudo mkdir /run/ATO
sudo chown $USER:$USER /run/ATO
chmod 755 /run/ATO
```

These last two steps (setting up the cgroup and `/run/ATO` directories) will need to be redone if you reboot.

I'm working on Dockerising this (#105), because it's admittedly rather involved.

Now run the server:

```bash
target/debug/attempt-this-online
```

This will bind to localhost on port 8500 by default. You can override this with the `ATO_BIND` environment variable.

### Tests
There are some basic sandbox functionality tests written in Python (make sure you have version 3.10 or higher installed).
The `test/run` helper script will set up the testing environment for you if needed (as well as running the tests).
Pass it the URL to the API, like this:

```bash
URL='ws://localhost:8500/api/v1/ws/execute' test/run
```

Some of the tests take a few seconds to run (because they test timing things). To skip these, set the `FAST` environment
variable:

```bash
FAST=1 URL='ws://localhost:8500/api/v1/ws/execute' test/run
```

### Automatic rebuilds
You may find it useful to have your code automatically rebuilt. Install [`entr`](https://eradman.com/entrproject/),
then run, in different shell instances:

```bash
ls src | entr -c cargo build --all-targets
```

```bash
export ATO_CGROUP_PATH=/sys/fs/cgroup/user.slice/user-$(id -u).slice/user@$(id -u).service/ATO
ls target/debug/attempt-this-online | entr -rc target/debug/attempt-this-online
```

To rerun tests automatically as well:

```bash
export URL='ws://localhost:8500/api/v1/ws/execute'
# export FAST=1
ls target/debug/attempt-this-online test/test.py | entr -c test/run
```

## Frontend developer instructions
See [`frontend/README.md`](../frontend/README.md).
