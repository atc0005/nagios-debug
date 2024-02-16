# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/nagios-debug/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.2.5] - 2024-02-16

### Changed

#### Dependency Updates

- (GH-274) canary: bump golang from 1.20.13 to 1.20.14 in /dependabot/docker/go
- (GH-261) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.3 to go-ci-oldstable-build-v0.14.4 in /dependabot/docker/builds
- (GH-267) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.4 to go-ci-oldstable-build-v0.14.5 in /dependabot/docker/builds
- (GH-268) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.5 to go-ci-oldstable-build-v0.14.6 in /dependabot/docker/builds
- (GH-277) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.6 to go-ci-oldstable-build-v0.14.9 in /dependabot/docker/builds
- (GH-279) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.9 to go-ci-oldstable-build-v0.15.0 in /dependabot/docker/builds
- (GH-266) go.mod: bump github.com/atc0005/go-nagios from 0.16.0 to 0.16.1

## [v0.2.4] - 2024-01-19

### Changed

#### Dependency Updates

- (GH-248) canary: bump golang from 1.20.11 to 1.20.12 in /dependabot/docker/go
- (GH-254) canary: bump golang from 1.20.12 to 1.20.13 in /dependabot/docker/go
- (GH-250) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.1 to go-ci-oldstable-build-v0.14.2 in /dependabot/docker/builds
- (GH-256) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.2 to go-ci-oldstable-build-v0.14.3 in /dependabot/docker/builds
- (GH-253) ghaw: bump github/codeql-action from 2 to 3

## [v0.2.3] - 2023-11-15

### Changed

#### Dependency Updates

- (GH-241) canary: bump golang from 1.20.10 to 1.20.11 in /dependabot/docker/go
- (GH-217) canary: bump golang from 1.20.7 to 1.20.8 in /dependabot/docker/go
- (GH-236) canary: bump golang from 1.20.8 to 1.20.10 in /dependabot/docker/go
- (GH-244) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.12 to go-ci-oldstable-build-v0.14.1 in /dependabot/docker/builds
- (GH-209) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.4 to go-ci-oldstable-build-v0.13.5 in /dependabot/docker/builds
- (GH-212) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.5 to go-ci-oldstable-build-v0.13.6 in /dependabot/docker/builds
- (GH-214) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.6 to go-ci-oldstable-build-v0.13.7 in /dependabot/docker/builds
- (GH-219) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.7 to go-ci-oldstable-build-v0.13.8 in /dependabot/docker/builds
- (GH-226) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.8 to go-ci-oldstable-build-v0.13.9 in /dependabot/docker/builds
- (GH-237) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.9 to go-ci-oldstable-build-v0.13.12 in /dependabot/docker/builds
- (GH-215) ghaw: bump actions/checkout from 3 to 4

## [v0.2.2] - 2023-08-18

### Added

- (GH-184) Add initial automated release notes config
- (GH-186) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.19.11` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.4` to `go-ci-oldstable-build-v0.13.4`
- (GH-188) Update Dependabot config to monitor both branches
- (GH-204) Update project to Go 1.20 series

## [v0.2.1] - 2023-07-14

### Overview

- RPM package improvements
- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.8` to `1.19.11`
  - `atc0005/go-nagios`
    - `v0.14.0` to `v0.16.0`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.4` to `go-ci-oldstable-build-v0.11.4`
- (GH-174) Update vuln analysis GHAW to remove on.push hook
- (GH-180) Update RPM postinstall scripts to use restorecon

### Fixed

- (GH-171) Disable depguard linter
- (GH-175) Restore local CodeQL workflow

## [v0.2.0] - 2023-04-06

### Overview

- Add support for generating DEB, RPM packages
- Build improvements
- Generated binary changes
  - filename patterns
  - compression (~ 66% smaller)
  - executable metadata
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-156) Generate RPM/DEB packages using nFPM
- (GH-157) Add version details to Windows executables

### Changed

- (GH-155) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-158) Makefile: Compress binaries & use fixed filenames
- (GH-159) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-160) Build dev/stable releases using go-ci Docker image

## [v0.1.18] - 2023-04-06

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-146) Add Go Module Validation, Dependency Updates jobs

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.8`
  - `atc0005/go-nagios`
    - `v0.10.2` to `v0.14.0`
- (GH-150) Drop `Push Validation` workflow
- (GH-151) Rework workflow scheduling
- (GH-153) Remove `Push Validation` workflow status badge

### Fixed

- (GH-162) Update vuln analysis GHAW to use on.push hook

## [v0.1.17] - 2022-12-09

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.1` to `1.19.4`
  - `atc0005/go-nagios`
    - `v0.10.0` to `v0.10.2`
- (GH-133) Refactor GitHub Actions workflows to import logic

### Fixed

- (GH-138) Fix Makefile Go module base path detection
- (GH-139) Replace literal/static newlines with pkg constant

## [v0.1.16] - 2022-09-22

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.1
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.9` to `1.19.1`
  - `atc0005/go-nagios`
    - `v0.8.2` to `v0.10.0`
  - `github/codeql-action`
    - `v2.1.22` to `v2.1.24`
- (GH-124) Update project to Go 1.19
- (GH-125) Update Makefile and GitHub Actions Workflows

### Fixed

- (GH-119) Update lintinstall Makefile recipe
- (GH-122) Add project doc.go file
- (GH-123) Add missing cmd doc file

## [v0.1.15] - 2022-05-06

### Overview

- Dependency updates
- built using Go 1.17.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.7` to `1.17.9`

## [v0.1.14] - 2022-03-03

### Overview

- Dependency updates
- CI / linting improvements
- built using Go 1.17.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.6` to `1.17.7`
  - `actions/checkout`
    - `v2.4.0` to `v3`
  - `actions/setup-node`
    - `v2.5.1` to `v3`

- (GH-104) Expand linting GitHub Actions Workflow to include `oldstable`,
  `unstable` container images
- (GH-105) Switch Docker image source from Docker Hub to GitHub Container
  Registry (GHCR)

## [v0.1.13] - 2022-01-25

### Overview

- Dependency updates
- built using Go 1.17.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.12` to `1.17.6`
    - (GH-100) Update go.mod file, canary Dockerfile to reflect current
      dependencies
  - `atc0005/go-nagios`
    - `v0.8.1` to `v0.8.2`

## [v0.1.12] - 2021-12-29

### Overview

- Dependency updates
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.10` to `1.16.12`
  - `actions/setup-node`
    - `v2.4.1` to `v2.5.1`

## [v0.1.11] - 2021-11-10

### Overview

- Dependency updates
- built using Go 1.16.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.8` to `1.16.10`
  - `atc0005/go-nagios`
    - `v0.7.0` to `v0.8.1`
  - `actions/checkout`
    - `v2.3.4` to `v2.4.0`
  - `actions/setup-node`
    - `v2.4.0` to `v2.4.1`

## [v0.1.10] - 2021-09-27

### Overview

- Dependency updates
- built using Go 1.16.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.7` to `1.16.8`
  - `atc0005/go-nagios`
    - `v0.6.1` to `v0.7.0`

## [v0.1.9] - 2021-08-09

### Overview

- Dependency updates
- built using Go 1.16.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.6` to `1.16.7`
  - `actions/setup-node`
    - updated from `v2.2.0` to `v2.4.0`

## [v0.1.8] - 2021-07-19

### Overview

- Dependency updates
- built using Go 1.16.6
  - **Statically linked**
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- Add "canary" Dockerfile to track stable Go releases, serve as a reminder to
  generate fresh binaries

### Changed

- dependencies
  - `Go`
    - `1.15.8` to `1.16.6`
  - `atc0005/go-nagios`
    - `v0.6.0` to `v0.6.1`
  - `actions/setup-node`
    - `v2.1.4` to `v2.2.0`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version

## [v0.1.7] - 2021-02-21

### Overview

- Dependency updates
- built using Go 1.15.8

### Changed

- dependencies
  - `go.mod` Go version
    - updated from `1.14` to `1.15`
  - built using Go 1.15.8
    - **Statically linked**
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `atc0005/go-nagios`
    - updated from `v0.5.2` to `v0.6.0`
  - `actions/setup-node`
    - updated from `v2.1.2` to `v2.1.4`

## [v0.1.6] - 2020-11-16

### Changed

- Dependencies
  - built using Go 1.15.5
    - **Statically linked**
    - Windows
      - x86
      - x64
    - Linux
      - x86
      - x64
  - `atc0005/go-nagios`
    - `v0.5.1` to `v0.5.2`
  - `actions/checkout`
    - `v2.3.3` to `v2.3.4`

## [v0.1.5] - 2020-10-11

### Added

- Binary release
  - Built using Go 1.15.2
  - **Statically linked**
  - Windows
    - x86
    - x64
  - Linux
    - x86
    - x64

### Changed

- Add `-trimpath` build flag

### Fixed

- Makefile build options do not generate static binaries

## [v0.1.4] - 2020-10-02

### Added

- Binary release
  - Built using Go 1.15.2
  - Windows
    - x86
    - x64
  - Linux
    - x86
    - x64

### Changed

- dependencies
  - `atc0005/go-nagios`
    - updated from `v0.3.0` to `v0.5.1`
  - `actions/setup-node`
    - updated from `v2.1.1` to `v2.1.2`
  - `actions/checkout`
    - updated from `v2.3.2` to `v2.3.3`

### Fixed

- Typo in README
- Makefile generates checksums with qualified path

## [v0.1.3] - 2020-08-22

### Added

- Docker-based GitHub Actions Workflows
  - Replace native GitHub Actions with containers created and managed through
    the `atc0005/go-ci` project.

  - New, primary workflow
    - with parallel linting, testing and building tasks
    - with three Go environments
      - "old stable"
      - "stable"
      - "unstable"
    - Makefile is *not* used in this workflow
    - staticcheck linting using latest stable version provided by the
      `atc0005/go-ci` containers

  - Separate Makefile-based linting and building workflow
    - intended to help ensure that local Makefile-based builds that are
      referenced in project README files continue to work as advertised until
      a better local tool can be discovered/explored further
    - use `golang:latest` container to allow for Makefile-based linting
      tooling installation testing since the `atc0005/go-ci` project provides
      containers with those tools already pre-installed
      - linting tasks use container-provided `golangci-lint` config file
        *except* for the Makefile-driven linting task which continues to use
        the repo-provided copy of the `golangci-lint` configuration file

  - Add Quick Validation workflow
    - run on every push, everything else on pull request updates
    - linting via `golangci-lint` only
    - testing
    - no builds

- Add new README badges for additional CI workflows
  - each badge also links to the associated workflow results

### Changed

- Disable `golangci-lint` default exclusions

- dependencies
  - `go.mod` Go version
    - updated from `1.13` to `1.14`
  - `actions/setup-go`
    - updated from `v2.1.0` to `v2.1.2`
      - since replaced with Docker containers
  - `actions/setup-node`
    - updated from `v2.1.0` to `v2.1.1`
  - `actions/checkout`
    - updated from `v2.3.1` to `v2.3.2`

- README
  - Link badges to applicable GitHub Actions workflows results

- Linting
  - Local
    - `Makefile`
      - install latest stable `golangci-lint` binary instead of using a fixed
          version
  - CI
    - remove repo-provided copy of `golangci-lint` config file at start of
      linting task in order to force use of Docker container-provided config
      file

## [v0.1.2] - 2020-07-06

### Changed

- Update dependencies
  - `actions/setup-go`
    - `v2.0.3` to `v2.1.0`
  - `actions/setup-node`
    - `v2.0.0` to `v2.1.0`
  - `atc0005/go-nagios`
    - `v0.2.0` to `v0.3.0`

- Rename binary from `debug` to `check_debug` to match naming pattern used by
  official plugins and our other Nagios plugins

### Fixed

- skip emission of `NAGIOS_LONGSERVICEOUTPUT` environment variable
  - this prevents a loop that otherwise occurs when the output from the plugin
    is captured and displayed again as `NAGIOS_LONGSERVICEOUTPUT` and ...

## [v0.1.1] - 2020-06-21

### Added

- Sort environment variables before listing them

### Fixed

- Newlines are incorrectly interpreted by Nagios

## [v0.1.0] - 2020-06-21

Initial release!

This release provides an early release version of a Nagios plugin used to
"echo" back detected environment variables and provided CLI arguments to help
troubleshoot other/"real" service checks.

### Added

- Initial Nagios plugin

- Go modules support (vs classic `GOPATH` setup)

[Unreleased]: https://github.com/atc0005/nagios-debug/compare/v0.2.5...HEAD
[v0.2.5]: https://github.com/atc0005/nagios-debug/releases/tag/v0.2.5
[v0.2.4]: https://github.com/atc0005/nagios-debug/releases/tag/v0.2.4
[v0.2.3]: https://github.com/atc0005/nagios-debug/releases/tag/v0.2.3
[v0.2.2]: https://github.com/atc0005/nagios-debug/releases/tag/v0.2.2
[v0.2.1]: https://github.com/atc0005/nagios-debug/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/nagios-debug/releases/tag/v0.2.0
[v0.1.18]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.18
[v0.1.17]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.17
[v0.1.16]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.16
[v0.1.15]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.15
[v0.1.14]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.14
[v0.1.13]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.13
[v0.1.12]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.12
[v0.1.11]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.11
[v0.1.10]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.10
[v0.1.9]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.9
[v0.1.8]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.8
[v0.1.7]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.7
[v0.1.6]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.6
[v0.1.5]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.5
[v0.1.4]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.4
[v0.1.3]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.0
