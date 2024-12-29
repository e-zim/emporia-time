# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][changelog], and this project adheres
to [Semantic Versioning][semver].

## [Unreleased]

### Maintenance

- Write reminders for the steps involved in wiki page updates
- Wait until tests pass before merging updates to dependencies
- Find missing references that the above changes might expect
- Check that upstream merges do not let dependencies skip test
- Setup the self hosted runner on other computed configuration

## [1.1.1] - 2024-09-08

### Added

- Release `README.md` and `CHANGELOG.md` and `LICENSE` file

### Fixed

- Include the manual pages as part of the release artifacts
- Adjust filepath separators to gather configured settings

### Maintenance

- Set the `go` version to a fixed `1.22.6` for the toolchain
- Update token authentication to use the latest Cognito SDK
- Test coverage as a metric to at least consider with changes
- Separate token exchanges into a cognito package for testing

## [1.1.0] - 2024-08-03

### Added

- Document expected behavior using the manual `man` pages

### Changed

- Replace relative paths in reference to a global command

## [1.0.2] - 2024-07-14

### Fixed

- Avoid assumptions about the $PATH to the `time` command

### Maintenance

- Enforce common linter checks before building the binary
- Format markdown files according to rules of the marksman
- Update go versions to 1.22 with other build dependencies
- Establish some standard setup to the self hosted runners
- Merge updates to dependencies after the test checks pass
- Build binaries for releases from an upstream nur package

## [1.0.1] - 2024-02-18

### Fixed

- Ignore temporary certificates for clean release signing

## [1.0.0] - 2024-02-18

### Added

- Print the current build version with a `--verison` flag

### Fixed

- Output outputs in an unbuffered manner as output happens
- Capture timing information for erroneous commands
- Sign and notarize packaged binaries made for macOS

### Maintenance

- Setup a Nix flake for more consistent developments
- Include a note to update milestones after a release
- Automatically extend the license at the start of a year
- Use the development Nix flake in continuous integration
- Checkout the entire Git history for automated testing
- Request that PR titles resemble a conventional commit

## [0.1.1] - 2024-01-15

### Added

- Package releases for many different operating systems

### Fixed

- Parse the provided help flag to display a helpful message
- Return the actual error that happens when building templates
- Avoid a panic when no command arguments are provided
- Output errors from parsing flags in a more clear manner

### Maintenance

- Perform asserts on uniquely named tests across packages
- Refactor command usage templating into a templates package
- Separate concerns of a single internal package into many
- Increment the end license year to include this new year
- Bump the Golang version to the most recent version of 1.21
- Include instructions for cutting and versioning a release
- Reduce frequency of dependabot updates to once a month
- Include dependency checks for actions in GitHub Actions

## [0.1.0] - 2023-07-08

### Added

- A helpful message is shown when the `--help` flags are used
- Login credentials can be provided with the `--username` and `--password` flags
- Environment variables `EMPORIA_USERNAME` and `EMPORIA_PASSWORD` will login too
- Specify a device to measure with flag `--device` or variable `EMPORIA_DEVICE`
- Detailed information about this program is included in the README
- Include the average power used over a command executation in watts
- Display measurements on separate lines with the `--portable` or `-p` flags
- Documentation created for collaboration and contribution processes

### Changed

- Bump Go version to 1.20
- Plainly use the /usr/bin/time command instead of a generic time
- Parse the `time` output for measurements represented as float64 seconds
- Display times with hours and minutes in the formatted output
- Match common output orderings with `-p` for portable parsing

### Fixed

- Clarify confusing or error prone steps in the getting started process
- Path to the repository was changed and now it matches
- Timing outputs are now consistent across operating systems
- Errors from the provided command are properly propogated
- Correctly dislpay energy as joules and power as watts
- Replace the command name in help templates on build errors

### Maintenance

- Check that changes are made to the changelog on changes
- Perform scheduled checks for dependency updates
- Measure energy usage as an integration test on the remote
- Perform the full authentication handshake on remote tests
- Restructure the repo to use multiple packages
- Prefer the .yml file extension in action workflows

## [0.0.2] - 2022-12-25

### Added

- Repeat lookups until a sureness of at least 80.0% is reached
- Check the status of Emporia's API before performing command
- Gather API tokens from AWS Cognito using Emporia credentials
- Select a device from connected Emporia devices
- Helpful information is displayed on an empty input
- Tests for energy conversions and usage extrapolation
- Linting checks and automated tests added on GitHub Actions

### Changed

- Execution info is output to stderr
- Config stored in `~/.config/etime/settings.json`
- Config file respects `XDG_CONFIG_HOME` environment variable
- Repository name changed to `emporia-time`
- Go version updated to 1.19

### Removed

- Tokens stored in environment variables
- Instructions for manually configuring devices

## [0.0.1] - 2022-11-27

### Added

- Perform user commands or scripts (without interactivity)
- Output results and timing information from user command
- Display energy usage stats from command duration
- Setup developer scripts in Makefile
- Instruct setup with a README
- Added the MIT license
- Created a CHANGELOG

<!-- a collection of links -->

[changelog]: https://keepachangelog.com/en/1.1.0/
[semver]: https://semver.org/spec/v2.0.0.html

<!-- a collection of releases -->

[Unreleased]: https://github.com/zimeg/emporia-time/compare/v1.1.1...HEAD
[1.1.1]: https://github.com/zimeg/emporia-time/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/zimeg/emporia-time/compare/v1.0.2...v1.1.0
[1.0.2]: https://github.com/zimeg/emporia-time/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/zimeg/emporia-time/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/zimeg/emporia-time/compare/v0.1.1...v1.0.0
[0.1.1]: https://github.com/zimeg/emporia-time/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/zimeg/emporia-time/compare/v0.0.2...v0.1.0
[0.0.2]: https://github.com/zimeg/emporia-time/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/zimeg/emporia-time/releases/tag/v0.0.1
