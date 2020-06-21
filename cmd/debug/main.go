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
	statusLineTmpl := "OK: %d environment variables, %d CLI arguments\n\n"

	numEnvVars := len(os.Environ())

	// first value is the path to the program, remaining values are the
	// arguments to the program
	numCLIArgs := len(os.Args) - 1

	fmt.Printf(
		statusLineTmpl,
		numEnvVars,
		numCLIArgs,
	)

	fmt.Printf("Environment variables:\n\n")

	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	fmt.Printf("\nCLI arguments:\n\n")

	switch {
	case numCLIArgs > 1:
		for num, arg := range os.Args[1:] {
			fmt.Printf(
				"Arg %d: %s\n",
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
