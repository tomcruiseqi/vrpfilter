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
)

func main() {
	output := flag.String("o", "result", "A file that contains the matches.")
	input := flag.String("i", "", "Source file.")
	others := flag.Args()

	// Parse the command line options.
	flag.Parse()

	others = append(others, *input)
	for _, file := range others {
		ok, err := Begin(file, *output)
		if !ok {
			fmt.Println(err)
		}
	}
}

// Open file and begin filtering.
func Begin(file string, result string) (ok bool, err error) {
	// Read the entire file to memory.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		ok = false
		return
	}

	// Convert to string.
	ok, err = grep(string(data), result)
	return
}

// Filter the text, find all of the license
// plate numbers and write to a file.
// If success, return ok as true. Otherwise, send an error to the caller
// and explain why fails.
func grep(data string, result string) (ok bool, err error) {
	ordinary, err := regexp.Compile(OrdinaryRe)
	if err != nil {
		ok = false
		return
	}

	armedpol, err := regexp.Compile(ArmedPolice)
	if err != nil {
		ok = false
		return
	}

	military, err := regexp.Compile(Military)
	if err != nil {
		ok = false
		return
	}

	police, err := regexp.Compile(Police)
	if err != nil {
		ok = false
		return
	}

	hm, err := regexp.Compile(HM)
	if err != nil {
		ok = false
		return
	}

	embassy, err := regexp.Compile(Embassy)
	if err != nil {
		ok = false
		return
	}

	// Filter the ordinary vehicle registration plate.
	orvrps := ordinary.FindAllString(data, -1)
	PrintVrps(orvrps)

	amvrps := armedpol.FindAllString(data, -1)
	PrintVrps(amvrps)

	milvrps := military.FindAllString(data, -1)
	PrintVrps(milvrps)

	polvrps := police.FindAllString(data, -1)
	PrintVrps(polvrps)

	hmvrps := hm.FindAllString(data, -1)
	PrintVrps(hmvrps)

	emvrps := embassy.FindAllString(data, -1)
	PrintVrps(emvrps)

	return true, nil
}

// Print the matches.
func PrintVrps(vrps []string) (bool, error) {
	if len(vrps) == 0 {
		return true, nil
	}

	fmt.Println(vrps)
	return true, nil
}
