<!-- omit in toc -->
# nagios-debug

Small Nagios plugin to return "debug" environment details

[![Latest Release](https://img.shields.io/github/release/atc0005/nagios-debug.svg?style=flat-square)](https://github.com/atc0005/nagios-debug/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/nagios-debug.svg)](https://pkg.go.dev/github.com/atc0005/nagios-debug)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/atc0005/nagios-debug)](https://github.com/atc0005/nagios-debug)
[![Lint and Build](https://github.com/atc0005/nagios-debug/actions/workflows/lint-and-build.yml/badge.svg)](https://github.com/atc0005/nagios-debug/actions/workflows/lint-and-build.yml)
[![Project Analysis](https://github.com/atc0005/nagios-debug/actions/workflows/project-analysis.yml/badge.svg)](https://github.com/atc0005/nagios-debug/actions/workflows/project-analysis.yml)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
- [Features](#features)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Installation](#installation)
  - [From source](#from-source)
  - [Using release binaries](#using-release-binaries)
- [Configuration options](#configuration-options)
- [Examples](#examples)
- [License](#license)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

Nagios plugin used to "echo" back detected environment variables and provided
CLI flags/values as part of service check output to help troubleshoot other
"real" service checks.

## Features

- single binary
- no configuration settings for the binary are required
- echoes back detected environment variables and provided CLI flags/values to
  help troubleshoot other service checks

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go
  - see this project's `go.mod` file for *preferred* version
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Windows 10
- Ubuntu Linux 18.04+

## Installation

### From source

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
   - NOTE: Pay special attention to the remarks about `$HOME/.profile`
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/nagios-debug`
   1. `cd nagios-debug`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     - `sudo yum install make gcc`
   - for Windows
     - Emulated environments (*easier*)
       - Skip all of this and build using the default `go build` command in
         Windows (see below for use of the `-mod=vendor` flag)
       - build using Windows Subsystem for Linux Ubuntu environment and just
         copy out the Windows binaries from that environment
       - If already running a Docker environment, use a container with the Go
         tool-chain already installed
       - If already familiar with LXD, create a container and follow the
         installation steps given previously to install required dependencies
     - Native tooling (*harder*)
       - see the StackOverflow Question `32127524` link in the
         [References](references.md) section for potential options for
         installing `make` on Windows
       - see the mingw-w64 project homepage link in the
         [References](references.md) section for options for installing `gcc`
         and related packages on Windows
1. Build binaries
   - for the current operating system, explicitly using bundled dependencies
         in top-level `vendor` folder
     - `go build -mod=vendor ./cmd/check_debug/`
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for use on Windows
      - `make windows`
   - for use on Linux
     - `make linux`
1. Copy the newly compiled binary from the applicable `/tmp` subdirectory path
   (based on the clone instructions in this section) below and deploy where
   needed.
   - if using `Makefile`
     - look in `/tmp/nagios-debug/check_debug/`
   - if using `go build`
     - look in `/tmp/nagios-debug/`

### Using release binaries

1. Download the [latest release][repo-url] binaries
1. Deploy
   - Place `check_debug` in the same location where your distro's package
     manager places other Nagios plugins
     - as `/usr/lib/nagios/plugins/check_whois` on Debian-based systems
     - as `/usr/lib64/nagios/plugins/check_whois` on RedHat-based systems

## Configuration options

None. Reference the generated binary from within your Nagios configuration as
you would another service check. The available environment variables and
provided CLI flags/values will be returned as part of the check output.

## Examples

TODO:

- example of creating a Nagios service check command
- example of referencing the newly created service check command in place of a
  real check
  - e.g., disk usage with custom thresholds
  - e.g., cert check with WARNING and CRITICAL expiration values

## License

From the [LICENSE](LICENSE) file:

```license
MIT License

Copyright (c) 2020 Adam Chalkley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## References

- <https://github.com/atc0005/go-nagios>
- <https://nagios-plugins.org/doc/guidelines.html>
- <https://gobyexample.com/environment-variables>
- <https://gobyexample.com/command-line-arguments>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/nagios-debug>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"

<!-- []: PLACEHOLDER "DESCRIPTION_HERE" -->
