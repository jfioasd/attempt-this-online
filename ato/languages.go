package ato

import "github.com/vmihailenco/msgpack/v5"

type language struct {
	Name     string `msgpack:"name"`
	Image    string `msgpack:"image"`
	Version  string `msgpack:"version"`
	Url      string `msgpack:"url"`
	Sbcs     bool   `msgpack:"sbcs"`
	SE_class string `msgpack:"SE_class"`
}

var serialisedLanguages []byte

func init() {
	b, err := msgpack.Marshal(Languages)
	if err != nil {
		panic(err)
	}
	serialisedLanguages = b
}

var Languages = map[string]language{
	"whython": {
		Name:     "Whython",
		Image:    "attemptthisonline/whython",
		Version:  "Latest",
		Url:      "https://github.com/pxeger/whython",
		Sbcs:     false,
		SE_class: "python",
	},
	"python": {
		Name:     "Python",
		Image:    "attemptthisonline/python_with_common_libraries",
		Version:  "Latest",
		Url:      "https://www.python.org",
		Sbcs:     false,
		SE_class: "python",
	},
	"zsh": {
		Name:     "Zsh",
		Image:    "attemptthisonline/zsh",
		Version:  "5",
		Url:      "https://www.zsh.org/",
		Sbcs:     false,
		SE_class: "bash",
	},
	"jelly": {
		Name:    "Jelly",
		Image:   "attemptthisonline/jelly",
		Version: "70c9fd93",
		Url:     "https://github.com/DennisMitchell/jellylanguage",
		Sbcs:    true,
	},
	"ruby": {
		Name:     "Ruby",
		Image:    "attemptthisonline/ruby",
		Version:  "Latest",
		Url:      "https://www.ruby-lang.org/",
		Sbcs:     false,
		SE_class: "ruby",
	},
	"python2": {
		Name:     "Python 2",
		Image:    "attemptthisonline/python2",
		Version:  "2",
		Url:      "https://docs.python.org/2/",
		Sbcs:     false,
		SE_class: "python2",
	},
	"scala3": {
		Name:    "Scala 3",
		Image:   "attemptthisonline/scala3",
		Version: "3",
		Url:     "https://www.scala-lang.org/",
		Sbcs:    false,
	},
	"scala2": {
		Name:    "Scala 2",
		Image:   "attemptthisonline/scala2",
		Version: "2",
		Url:     "https://www.scala-lang.org/",
		Sbcs:    false,
	},
	"java": {
		Name:     "Java",
		Image:    "attemptthisonline/java",
		Version:  "Latest",
		Url:      "https://en.wikipedia.org/wiki/Java_(programming_language)",
		Sbcs:     false,
		SE_class: "java",
	},
	"tictac": {
		Name:    "Tictac",
		Image:   "attemptthisonline/tictac",
		Version: "Latest",
		Url:     "https://github.com/pxeger/tictac",
		Sbcs:    true,
	},
	"bash": {
		Name:    "Bash",
		Image:   "attemptthisonline/bash",
		Version: "Latest",
		Url:     "https://www.gnu.org/software/bash/",
		Sbcs:    false,
	},
	"pip": {
		Name:    "Pip",
		Image:   "attemptthisonline/pip",
		Version: "Latest",
		Url:     "https://github.com/dloscutoff/pip",
		Sbcs:    false,
	},
	"funky2": {
		Name:    "Funky2",
		Image:   "attemptthisonline/funky2",
		Version: "Latest",
		Url:     "https://funky2.a-ta.co/",
		Sbcs:    false,
	},
	"c_gcc": {
		Name:    "C (GCC)",
		Image:   "attemptthisonline/base",
		Version: "11",
		Url:     "https://gcc.gnu.org",
		Sbcs:    false,
	},
	"cplusplus_gcc": {
		Name:    "C++ (GCC)",
		Image:   "attemptthisonline/gcc",
		Version: "11",
		Url:     "https://gcc.gnu.org",
		Sbcs:    false,
	},
	"objective_cplusplus_gcc": {
		Name:    "Objective-C++ (GCC)",
		Image:   "attemptthisonline/gcc",
		Version: "11",
		Url:     "https://gcc.gnu.org",
		Sbcs:    false,
	},
	"objective_c_gcc": {
		Name:    "Objective-C (GCC)",
		Image:   "attemptthisonline/gcc",
		Version: "11",
		Url:     "https://gcc.gnu.org",
		Sbcs:    false,
	},
	"go_gcc": {
		Name:    "Go (GCC)",
		Image:   "attemptthisonline/gcc",
		Version: "11",
		Url:     "https://gcc.gnu.org",
		Sbcs:    false,
	},
	"gnat": {
		Name:    "Ada (GNAT)",
		Image:   "attemptthisonline/gcc",
		Version: "11",
		Url:     "https://en.wikipedia.org/wiki/GNAT",
		Sbcs:    false,
	},
	"gfortran": {
		Name:    "Fortran (GFortran)",
		Image:   "attemptthisonline/gcc",
		Version: "11",
		Url:     "https://gcc.gnu.org/fortran",
		Sbcs:    false,
	},
	"gdc": {
		Name:    "D (GDC)",
		Image:   "attemptthisonline/gcc",
		Version: "11",
		Url:     "https://gdcproject.org",
		Sbcs:    false,
	},
	"node": {
		Name:    "JavaScript (Node.js)",
		Image:   "attemptthisonline/node",
		Version: "Latest",
		Url:     "https://nodejs.org",
		Sbcs:    false,
	},
	"vyxal": {
		Name:    "Vyxal",
		Image:   "attemptthisonline/vyxal",
		Version: "2",
		Url:     "https://github.com/Vyxal/Vyxal",
		Sbcs:    false,
	},
	"go": {
		Name:    "Go",
		Image:   "attemptthisonline/go",
		Version: "Latest",
		Url:     "https://go.dev",
		Sbcs:    false,
	},
	"perl": {
		Name:    "Perl",
		Image:   "attemptthisonline/perl",
		Version: "Latest",
		Url:     "https://www.perl.org",
		Sbcs:    false,
	},
	"php": {
		Name:    "PHP",
		Image:   "attemptthisonline/php",
		Version: "Latest",
		Url:     "https://www.php.net",
		Sbcs:    false,
	},
	"r": {
		Name:    "R",
		Image:   "attemptthisonline/r_but_longer",
		Version: "Latest",
		Url:     "https://www.r-project.org",
		Sbcs:    false,
	},
	"erlang": {
		Name:    "Erlang",
		Image:   "attemptthisonline/erlang",
		Version: "Latest",
		Url:     "https://www.erlang.org",
		Sbcs:    false,
	},
	"elixir": {
		Name:    "Elixir",
		Image:   "attemptthisonline/elixir",
		Version: "Latest",
		Url:     "https://elixir-lang.org",
		Sbcs:    false,
	},
	"guile": {
		Name:    "Guile",
		Image:   "attemptthisonline/base",
		Version: "Latest",
		Url:     "https://www.gnu.org/software/guile/",
		Sbcs:    false,
	},
	"julia": {
		Name:    "Julia",
		Image:   "attemptthisonline/julia",
		Version: "Latest",
		Url:     "https://julialang.org",
		Sbcs:    false,
	},
	"deno": {
		Name:    "TypeScript (Deno)",
		Image:   "attemptthisonline/deno",
		Version: "Latest",
		Url:     "https://deno.land",
		Sbcs:    false,
	},
	"kotlin": {
		Name:    "Kotlin",
		Image:   "attemptthisonline/kotlin",
		Version: "Latest",
		Url:     "https://kotlinlang.org",
		Sbcs:    false,
	},
	"rust": {
		Name:    "Rust",
		Image:   "attemptthisonline/rust",
		Version: "Latest",
		Url:     "https://www.rust-lang.org",
		Sbcs:    false,
	},
	"clang": {
		Name:    "C (clang)",
		Image:   "attemptthisonline/clang",
		Version: "Latest",
		Url:     "https://clang.llvm.org",
		Sbcs:    false,
	},
	"k_ok": {
		Name:    "K (oK)",
		Image:   "attemptthisonline/ok",
		Version: "Latest",
		Url:     "https://github.com/JohnEarnest/ok",
		Sbcs:    false,
	},
	"haskell": {
		Name:    "Haskell",
		Image:   "attemptthisonline/haskell",
		Version: "Latest",
		Url:     "https://www.haskell.org",
		Sbcs:    false,
	},
	"quipu": {
		Name:    "Quipu",
		Image:   "attemptthisonline/quipu",
		Version: "Latest",
		Url:     "https://github.com/cgccuser/quipu",
		Sbcs:    false,
	},
	"brainfuck": {
		Name:    "Brainfuck",
		Image:   "attemptthisonline/tio_brainfuck",
		Version: "the only version",
		Url:     "https://github.com/TryItOnline/brainfuck",
		Sbcs:    false,
	},
	"alphuck": {
		Name:    "Alphuck",
		Image:   "attemptthisonline/tio_brainfuck",
		Version: "Unknown",
		Url:     "https://github.com/TryItOnline/brainfuck",
		Sbcs:    false,
	},
	"brainbool": {
		Name:    "Brainbool",
		Image:   "attemptthisonline/tio_brainfuck",
		Version: "Unknown",
		Url:     "https://github.com/TryItOnline/brainfuck",
		Sbcs:    false,
	},
	"ooocode": {
		Name:    "oOo CODE",
		Image:   "attemptthisonline/tio_brainfuck",
		Version: "Unknown",
		Url:     "https://github.com/TryItOnline/brainfuck",
		Sbcs:    false,
	},
	"tinybf": {
		Name:    "TinyBF",
		Image:   "attemptthisonline/tio_brainfuck",
		Version: "Unknown",
		Url:     "https://github.com/TryItOnline/brainfuck",
		Sbcs:    false,
	},
	"random_brainfuck": {
		Name:    "Random Brainfuck",
		Image:   "attemptthisonline/tio_brainfuck",
		Version: "Unknown",
		Url:     "https://github.com/TryItOnline/brainfuck",
		Sbcs:    false,
	},
	"brainlove": {
		Name:    "Brainlove",
		Image:   "attemptthisonline/tio_brainfuck",
		Version: "Unknown",
		Url:     "https://github.com/TryItOnline/brainfuck",
		Sbcs:    false,
	},
	"extended_brainfuck_type_I": {
		Name:    "Extended Brainfuck Type I",
		Image:   "attemptthisonline/tio_brainfuck",
		Version: "Unknown",
		Url:     "https://github.com/TryItOnline/brainfuck",
		Sbcs:    false,
	},
	"bqn": {
		Name:    "BQN (CBQN)",
		Image:   "attemptthisonline/cbqn",
		Version: "Latest",
		Url:     "https://mlochbaum.github.io/BQN/",
		Sbcs:    false,
	},
	/* "ngn_apl": {
		Name:    "APL (ngn/apl)",
		Image:   "attemptthisonline/ngn_apl",
		Version: "Latest",
		Url:     "https://github.com/abrudz/ngn-apl",
		Sbcs:    false,
	}, */
	"factor": {
		Name:    "Factor",
		Image:   "attemptthisonline/factor",
		Version: "Latest",
		Url:     "https://factorcode.org",
		Sbcs:    false,
	},
	"sbcl": {
		Name:    "Common Lisp (SBCL)",
		Image:   "attemptthisonline/sbcl",
		Version: "Latest",
		Url:     "http://www.sbcl.org",
		Sbcs:    false,
	},
	"k_ngn": {
		Name:    "K (ngn/k)",
		Image:   "attemptthisonline/ngn_k",
		Version: "Latest",
		Url:     "https://codeberg.org/ngn/k",
		Sbcs:    false,
	},
	"apl_dzaima": {
		Name:    "APL (dzaima/APL)",
		Image:   "attemptthisonline/dzaima_apl",
		Version: "Latest",
		Url:     "https://github.com/dzaima/APL",
		Sbcs:    true,
	},
	"lua": {
		Name:    "Lua",
		Image:   "attemptthisonline/lua",
		Version: "Latest",
		Url:     "https://www.lua.org",
		Sbcs:    false,
	},
	"crystal": {
		Name:    "Crystal",
		Image:   "attemptthisonline/crystal",
		Version: "Latest",
		Url:     "https://crystal-lang.org/",
		Sbcs:    false,
	},
	"nim": {
		Name:    "Nim",
		Image:   "attemptthisonline/nim",
		Version: "Latest",
		Url:     "https://nim-lang.org/",
		Sbcs:    false,
	},
	"neko": {
		Name:    "Neko",
		Image:   "attemptthisonline/neko",
		Version: "Latest",
		Url:     "https://nekovm.org/",
		Sbcs:    false,
	},
	"zig": {
		Name:    "Zig",
		Image:   "attemptthisonline/zig",
		Version: "Latest",
		Url:     "https://ziglang.org/",
		Sbcs:    false,
	},
	"slashes": {
		Name:    "///",
		Image:   "attemptthisonline/slashes",
		Version: "Latest",
		Url:     "https://esolangs.org/wiki////",
		Sbcs:    false,
	},
	"sed": {
		Name:    "sed",
		Image:   "attemptthisonline/base",
		Version: "Latest",
		Url:     "https://www.gnu.org/software/sed/",
		Sbcs:    false,
	},
	"awk": {
		Name:    "AWK",
		Image:   "attemptthisonline/base",
		Version: "Latest",
		Url:     "https://www.gnu.org/software/gawk/",
		Sbcs:    false,
	},
	"jq": {
		Name:    "jq",
		Image:   "attemptthisonline/jq",
		Version: "Latest",
		Url:     "https://stedolan.github.io/jq/",
		Sbcs:    false,
	},
	"yq": {
		Name:    "yq",
		Image:   "attemptthisonline/jq",
		Version: "Latest",
		Url:     "https://kislyuk.github.io/yq/",
		Sbcs:    false,
	},
	"bc": {
		Name:    "bc",
		Image:   "attemptthisonline/bc",
		Version: "Latest",
		Url:     "https://www.gnu.org/software/bc/",
		Sbcs:    false,
	},
	"dc": {
		Name:    "dc",
		Image:   "attemptthisonline/bc",
		Version: "Latest",
		Url:     "https://www.gnu.org/software/bc/",
		Sbcs:    false,
	},
	"tcl": {
		Name:    "Tcl",
		Image:   "attemptthisonline/tcl",
		Version: "Latest",
		Url:     "http://tcl.sourceforge.net",
		Sbcs:    false,
	},
	"j": {
		Name:    "J",
		Image:   "attemptthisonline/j_but_longer",
		Version: "Latest",
		Url:     "https://www.jsoftware.com",
		Sbcs:    false,
	},
	"pari_gp": {
		Name:    "PARI/GP",
		Image:   "attemptthisonline/pari_gp",
		Version: "Latest",
		Url:     "https://pari.math.u-bordeaux.fr",
		Sbcs:    false,
	},
	"exceptionally": {
		Name:    "Exceptionally",
		Image:   "attemptthisonline/exceptionally",
		Version: "Latest",
		Url:     "https://github.com/dloscutoff/Esolangs/tree/master/Exceptionally",
		Sbcs:    false,
	},
	"regenerate": {
		Name:    "Regenerate",
		Image:   "attemptthisonline/regenerate",
		Version: "Latest",
		Url:     "https://github.com/dloscutoff/Esolangs/tree/master/Regenerate",
		Sbcs:    false,
	},
	"dirac": {
		Name:    "dirac",
		Image:   "attemptthisonline/dirac",
		Version: "Latest",
		Url:     "https://esolangs.org/wiki/Dirac",
		Sbcs:    false,
	},
	"j_uby": {
		Name:    "J-uby",
		Image:   "attemptthisonline/j_uby",
		Version: "Latest",
		Url:     "https://github.com/cyoce/J-uby",
		Sbcs:    false,
	},
	"dyalog_apl": {
		Name:    "APL (Dyalog APL)",
		Image:   "dyalog/dyalog",
		Version: "Latest",
		Url:     "https://www.dyalog.com/products.htm",
		Sbcs:    false,
	},
	"05ab1e": {
		Name:    "05AB1E",
		Image:   "attemptthisonline/05ab1e",
		Version: "Latest",
		Url:     "https://github.com/Adriandmen/05AB1E",
		Sbcs:    true,
	},
	"tex": {
		Name:     "TeX",
		Image:    "attemptthisonline/texlive",
		Version:  "Latest",
		Url:      "https://tug.org/texlive/doc.html",
		Sbcs:     false,
		SE_class: "tex",
	},
	"flax": {
		Name:    "flax",
		Image:   "attemptthisonline/flax",
		Version: "Latest",
		Url:     "https://github.com/PyGamer0/flax",
		Sbcs:    true,
	},
	"whitespace": {
		Name:    "Whitespace",
		Image:   "attemptthisonline/whitespace",
		Version: "Latest",
		Url:     "https://web.archive.org/web/20150618184706/http://compsoc.dur.ac.uk/whitespace/tutorial.php",
		Sbcs:    false,
	},
	"charcoal": {
		Name:    "Charcoal",
		Image:   "attemptthisonline/charcoal",
		Version: "Latest",
		Url:     "https://github.com/somebody1234/Charcoal",
		Sbcs:    true,
	},
	"husk": {
		Name:    "Husk",
		Image:   "attemptthisonline/husk",
		Version: "Latest",
		Url:     "https://github.com/barbuz/Husk",
		Sbcs:    true,
	},
	"nibbles": {
		Name:    "Nibbles",
		Image:   "attemptthisonline/nibbles",
		Version: "Latest",
        Url:     "http://golfscript.com/nibbles/index.html",
		Sbcs:    false,
	},
}
