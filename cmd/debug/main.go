// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/nagios-debug
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"os"

	"github.com/atc0005/go-nagios"
)

func main() {

	// prepare initial status line
	statusLineTmpl := "OK: %d environment variables, %d CLI arguments\r\n\r\n"

	numEnvVars := len(os.Environ())

	// first value is the path to the program, remaining values are the
	// arguments to the program
	numCLIArgs := len(os.Args) - 1

	fmt.Printf(
		statusLineTmpl,
		numEnvVars,
		numCLIArgs,
	)

	fmt.Printf("Environment variables:\r\n\r\n")

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	fmt.Printf("\r\nCLI arguments:\r\n\r\n")

	switch {
	case numCLIArgs > 1:
		for num, arg := range os.Args[1:] {
			fmt.Printf(
				"Arg %d: %s\r\n",
				num,
				arg,
			)
		}
	default:
		fmt.Println("No CLI flags defined")
	}

	// Assume all is well for now
	os.Exit(nagios.StateOK)
}
