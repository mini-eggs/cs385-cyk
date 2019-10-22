package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Grammar data types.
// Defines the structure we're willing to accept from json files.

type Grammar struct {
	Productions []Production
}

type Production struct {
	Left  string
	Right [][]string
}

// Helper functions.

// matrixfromstring - given a string break it into
// a matrix in a specific way.
// example:
//   input: "abcd"
//   output: [ [ "a", "bcd" ],
//             [ "ab", "cd" ],
//             [ "abc", "d" ] ]
func matrixfromstring(w string) [][]string {
	l := len(w)

	if l < 2 {
		return [][]string{[]string{w}}
	}

	ret := [][]string{}

	for i := 1; i < l; i++ {
		ret = append(ret, []string{
			w[0:i],
			w[i:l],
		})
	}

	return ret
}

// matrixmerge - merge arrays in a specific way
// example:
// input:
//   [ [ A, B, C ],
//     [ M, N ],
//     [ X, Y ] ]
// output:
//   [ [ A M X ],
//   [ [ A M Y ],
//   [ [ A N X ],
//   [ [ A N Y ],
//   [ [ B M X ],
//   [ [ B M Y ],
//   [ [ B N X ],
//   [ [ B N Y ],
//   [ [ C M X ],
//   [ [ C M Y ],
//   [ [ C N X ],
//   [ [ C N Y ] ]
func matrixmerge(in [][]string) (out [][]string) {
	// base case
	if len(in) < 2 {
		for _, set := range in {
			for _, value := range set {
				out = append(out, []string{value})
			}
		}
		return
	}

	// all other cases

	first := in[0]
	rest := matrixmerge(in[1:])

	for _, firstItem := range first {
		for _, row := range rest {
			out = append(out, append([]string{firstItem}, row...))
		}
	}

	return out
}

// cubemerge - one level above cube merge
// take a n^3 array and merge it down into n^2
func cubemerge(cube [][][]string) [][]string {
	ret := [][]string{}
	for _, matrix := range cube {
		ret = append(ret, matrixmerge(matrix)...)
	}
	return ret
}

// matrixcompare - given two simple arrays
// check if they are equal
func matrixcompare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if strings.Compare(a[i], b[i]) != 0 {
			return false
		}
	}

	return true
}

// CYK algorithm.
// Given a grammar in CNF and a string determine if the there is or is not a
// derivation for the string.
// https://en.wikipedia.org/wiki/CYK_algorithm
func CYK(w string, g Grammar) (bool, error) {
	if len(g.Productions) < 1 {
		return false, errors.New("grammar has no productions; must have at least one production")
	}

	// The start symbol is assumed to be the
	// variable on the left side of our first
	// production.
	startSym := g.Productions[0].Left

	// Keep track of what generates what
	// i.e. our V00 - Vxy values.
	hash := make(map[string][]string)

	// Length of input, just so we don't have to continually recalc it.
	inputLen := len(w)

	// Initate algo with O^2 loop for the purpose of iterating through strings
	// like so:
	// if given w is "abc" we will loop and set the value of strToMatch for:
	//   "a", "b", "c", "ab", "bc", "abc"
	// for a total of six loops.

	for outer := 0; outer < inputLen; outer++ {

		// how many letters we're matching against for
		// current outer loop
		// ie first one char, then two, ..., up to inputLen
		letterCountToMatch := outer + 1

		for inner := 0; inner < inputLen-outer; inner++ {

			strToMatch := w[inner : inner+letterCountToMatch]

			// base case
			// our initial hash is empty. Fill hash up with productions that produce
			// terminals.

			if outer == 0 { /* first loop */
				for _, row := range g.Productions {
					for _, prod := range row.Right {
						fullStr := strings.Join(prod, "")
						if strToMatch == fullStr {
							hash[fullStr] = append(hash[fullStr], row.Left)
						}
					}
				}
			} else { /* all other loops */

				// regular case

				// break up our current input by:
				//   for an input "abcd"
				//   our matrix will be:
				//     [ a, bcd ,
				//     , ab, cd ,
				//     , adc, d ]
				matrices := [][][]string{}

				for _, stringSet := range matrixfromstring(strToMatch) {
					matrice := [][]string{}
					for _, str := range stringSet {
						matrice = append(matrice, hash[str])
					}
					matrices = append(matrices, matrice)
				}

				// Perform matrix multiplication on our matrices from above. For each
				// result check if the values can be obtained from production rules in
				// the grammar g. If so state so by storing the given left hand side of
				// the production in our hash, using a key of the current string,
				// strToMatch.
				for _, matrix := range cubemerge(matrices) {
					for _, row := range g.Productions {
						for _, prod := range row.Right {
							if matrixcompare(matrix, prod) {
								hash[strToMatch] = append(hash[strToMatch], row.Left)
							}
						}
					}
				}
			}
		}
	}

	// Loop through all keys of hash. For all keys of the same length of our input
	// check if we have a start symbol in the given key's values. If so return
	// true, the grammar can generate the string.
	for key := range hash {
		if len(key) == inputLen {
			for _, value := range hash[key] {
				if value == startSym {
					return true, nil
				}
			}
		}
	}

	// We failed to find a start symbol in our final iterations hash values. The
	// grammar can't generate the input!
	return false, nil
}

// Entry to program. Read command line arguments to read a grammar from a json
// file. Then take input from user for what string user is interested in and
// state whether or not the grammar can produce the string.
func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("must enter command line arguement of json file")
		return
	}

	for index, filename := range args {
		// Skip first case, this is the program name do to
		// how command line args work, i.e.
		// $ ./out/cyk-linux-bin %FILE
		// if index == 0 then filename == "./out/cyk-linux-bin"
		// obviously not a file!
		if index == 0 {
			continue
		}

		// Read file and report errors.
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
			return
		}

		var grammar Grammar

		// Attempt to wrangle file into grammar.
		// Report errors.
		err = json.Unmarshal(bytes, &grammar)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Print("enter input text: ")

		// Take input from terminal.
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}

		// Remove whitespace from user input.
		input = strings.TrimSpace(input)

		// Work!
		val, err := CYK(input, grammar)
		if err != nil {
			log.Fatal(err)
			return
		}

		if val {
			fmt.Printf("%s is generated by the grammar\n", input)
		} else {
			fmt.Printf("%s is not generated by the grammar\n", input)
		}
	}
}
