/*
Package vrpfilter implments to find all of the Vehicle Registration Plates
contained in a file.

Usage of this tool:

		vrpfilter -o result -i file1 [file2] [...]

		result is a file which holds all the matches, the default is
		result.

		the inputs (-i) must be specified by at lease one argument
		which can be absolute or relative.
*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	output := flag.String("o", "result", "A file that contains the matches.")
	input := flag.String("i", "file1", "Source file.")
	others := flag.Args()

	// Parse the command line options.
	flag.Parse()
}

// Filter the text inputed as stream, find all of the license plate numbers
// and write to a file.
// If success, return ok as true. Otherwise, send an error to the caller
// and explain why fails.
func grep(reader io.Reader, result *string) (ok bool, err error) {
}
