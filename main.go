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
	"io/ioutil"
	"regexp"
)

const (
	OrdinaryRe  = `[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼][A-Z][A-HJ-NP-Z0-9]{5}`
	ArmedPolice = `WJ[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼]?[0-9]{4}[A-Z0-9]`
	Military    = `[A-Z]{2}[0-9]{5}`
	Police      = `[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼][A-Z0-9][0-9]{3}警`
	HM          = `粤Z[A-Z0-9]{4}[港澳]`
	Embassy     = `使[0-9]{6}`
	ReCount     = 6
)

func main() {
	output := flag.String("o", "result", "A file that contains the matches.")
	input := flag.String("i", "", "Source file.")
	others := flag.Args()

	// Parse the command line options.
	flag.Parse()

	others = append(others, *input)
	for _, file := range others {
		err := Begin(file, *output)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Open file and begin filtering.
// Save the mateches to the file specified by the second parameter.
func Begin(file string, result string) error {
	// Read the entire file to memory.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// Begin filtering.
	// At present, only print the matches to screen.
	return grep(string(data))
}

// Filter the text, find all of the license
// plate numbers and write to screen.
//  Otherwise, send an error to the caller
func grep(data string) error {
	matches := make(chan *[]string, 100)
	go findvrps(OrdinaryRe, data, matches)
	go findvrps(ArmedPolice, data, matches)
	go findvrps(Military, data, matches)
	go findvrps(Police, data, matches)
	go findvrps(HM, data, matches)
	go findvrps(Embassy, data, matches)
	printvrps(matches)
	return nil
}

// Apply regular expression to the data given.
func findvrps(regex string, data string, matches chan *[]string) {
	regobj, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println("error occurs. regex - [%v] data - [%v] error - [%v]", regex, data, err)
		return
	}

	// Begin filtering.
	mares := regobj.FindAllString(data, -1)
	matches <- &mares
	matches <- nil
}

// Print the matches.
func printvrps(matches chan *[]string) {
	finished := 0
	for match := range matches {
		// One kind of match has finished.
		if match == nil {
			finished++

			if ReCount == finished {
				return
			}
			continue
		}

		fmt.Println(*match)
	}
}
