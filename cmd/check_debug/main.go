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
	"sort"
	"strings"

	"github.com/atc0005/go-nagios"
)

func main() {

	const NagiosLongServiceOutputEnvVar string = "NAGIOS_LONGSERVICEOUTPUT"

	// prepare initial status line
	statusLineTmpl := "OK: %d environment variables, %d CLI arguments%[3]s%[3]s"

	numEnvVars := len(os.Environ())

	// first value is the path to the program, remaining values are the
	// arguments to the program
	numCLIArgs := len(os.Args) - 1

	fmt.Printf(
		statusLineTmpl,
		numEnvVars,
		numCLIArgs,
		nagios.CheckOutputEOL,
	)

	fmt.Printf("Environment variables:%[1]s%[1]s", nagios.CheckOutputEOL)
	fmt.Printf(
		"NOTE: Skipping emission of %s to help prevent a loop.%[2]s%[2]s",
		NagiosLongServiceOutputEnvVar,
		nagios.CheckOutputEOL,
	)

	origEnvVars := os.Environ()
	sortedEnvVars := make([]string, len(origEnvVars))
	copy(sortedEnvVars, origEnvVars)

	sort.Slice(sortedEnvVars, func(i, j int) bool {
		return sortedEnvVars[i] < sortedEnvVars[j]
	})

	for _, e := range sortedEnvVars {
		// skip emission of NAGIOS_LONGSERVICEOUTPUT to prevent a loop
		if strings.HasPrefix(
			strings.ToUpper(e),
			NagiosLongServiceOutputEnvVar,
		) {
			continue
		}
		fmt.Printf("%s%s", e, nagios.CheckOutputEOL)
	}

	fmt.Printf("%[1]sCLI arguments:%[1]s%[1]s", nagios.CheckOutputEOL)

	switch {
	case numCLIArgs > 1:
		for num, arg := range os.Args[1:] {
			fmt.Printf(
				"Arg %d: %s%[3]s",
				num,
				arg,
				nagios.CheckOutputEOL,
			)
		}
	default:
		fmt.Println("No CLI flags defined")
	}

	// Assume all is well for now
	os.Exit(nagios.StateOKExitCode)
}
