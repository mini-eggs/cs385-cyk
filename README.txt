ABOUT

This program is an extra credit assigment for CS385 at the University of Idaho.
It's an implementation of the CYK algorithm in the Go programming language. The
program contains tests and example inputs to test the validity of the program.

HOW TO RUN

On the release page of this Github you will find x86_64 binaries for Windows,
MacOS, and GNU/Linux. Once you have downloaded the program you will need to
execute it on the command line while providing at least one filepath to a json
file that will represent a grammar, like: 

$ ./cyk-linux-bin example_input_one.json

From there the program will read the json files (and let you know if it's
malformed) then prompt the user for an input string. "bbabb" is a good one. The
program will then execute the CYK function in the main.go files and detemine if
there is a possible derivation for the input string or not.

TESTING

If you have the Go programming language installed, running the tests (within
main_test.go -- there's only three of them) is a simple:

$ go test

BUILDING 

After you have cloned the repository building a "cyk" binary is a simple: 

$ go build

assuming you're in the root of the repository.

Enjoy.


