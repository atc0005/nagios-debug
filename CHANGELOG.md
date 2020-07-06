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

[Unreleased]: https://github.com/atc0005/nagios-debug/compare/v0.1.2...HEAD
[v0.1.2]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/nagios-debug/releases/tag/v0.1.0
