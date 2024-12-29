# Maintainers guide

Hey there! It's about time... Watt have you been jouling!?

**Outline**:

- [Project setup](#project-setup)
- [Testing](#testing)
- [Updating the wiki](#updating-the-wiki)
- [Merging pull requests](#merging-pull-requests)
- [Cutting a release](#cutting-a-release)

## Project setup

Building from source to reflect any code changes only takes a few fast steps.

1. Install the latest version of [Go][golang].
2. From a directory for development, download the source and compile `etime`:

```sh
$ git clone https://github.com/zimeg/emporia-time.git
$ cd emporia-time
$ make build
```

An [understanding of Go][learn_go] is a likely prerequisite for any programming
and can be an enjoyable language to learn!

### Nix configuration

A prepared development environment can be guaranteed from the `flake.nix`:

```sh
$ nix develop
```

Using [Nix][nix] is completely optional but somewhat recommended for
consistency.

### Project structure

This project hopes to use different directories to separate various concerns,
currently using the following structure:

- `/` – primary project files and metadata for the repository
- `.github/` – information for collaboration and continuous integrations
- `cmd/` – controllers for the different stages of the command
- `internal/` – helpful utilities needed to create the program
- `pkg/` – various concerns that are pieced together to form the program

### Makefile commands

For ease of development, some commands are added in a `Makefile`:

- `make build` – build the program binary
- `make staging` – package a distribution
- `make release` – sign and notarize packages
- `make test` – perform the written code tests
- `make clean` – remove all program artifacts

The name of the binary can be changed with `make build BIN=timer`.

## Testing

All tests should aim to cover common cases, both happy and erroneous, to build
confidence in any changes.

### Unit tests

Written tests should reside in a file adjacent to the functionality being tested
and suffixed with `_test.go`.

All tests can be run with `make test` and example test cases can be found
throughout the repo.

While coverage isn't critical, various permutations of input are often used to
check edge cases. There's some balance.

### Integration tests

Assurance that the program works as expected with the Emporia API can be gained
by running the program with any command:

```sh
$ make build
$ ./etime sleep 4
```

A smart plug and Emporia credentials are needed for this to be successful.

### On the remote

When changes are proposed or made to the remote repository, the full test suite
is performed to verify stability in any changes.

A [**new self-hosted runner**][runner] with a connected device can be brought
online to test changes with custom values for variables:

- `EMPORIA_DEVICE`
- `EMPORIA_USERNAME`
- `EMPORIA_PASSWORD`

Additionally, some change to the `CHANGELOG.md` is checked for on pull requests.

## Updating the wiki

Occasional reminders from PG&E are sent with details for the [wiki][wiki] pages.
Some discretion around sharing account numbers is recommended, and liberties are
encouraged for emoji replacement, but numbers for pricing ought to remain right.

### Writing a file

Markdown files that make these pages exist in a hidden repo that moves all files
to the top level path:

```sh
$ git clone https://github.com/zimeg/emporia-time.wiki.git
$ cd emporia-time.wiki
$ vim -o _Sidebar.md Statements/March-2080.md
$ git commit --all -m "chore: upload the statement from 2080-03-08 as markdown"
$ git push
$ open https://github.com/zimeg/emporia-time/wiki/March-2080
```

Emails with uploaded details should be archived or deleted or removed from thou
inbox once complete.

## Merging pull requests

Confidence in the tests should cover edge cases well enough to trust the suite.
A green status signals nothing broke as a result of changes, and an example run
can be seen in the actions output.

On any change, the following should be verified before merging:

- Documentation is correct and updated everywhere necessary
- Code changes move the project in a positive direction

If that all looks good and the change is solid, the **Squash and merge** awaits.

## Cutting a release

When the time is right to bump versions, either for new features or bug fixes,
the following steps can be taken:

1. Add the new version header to the `CHANGELOG.md` to mark the release
2. Preemptively update the version links at the end of the `CHANGELOG.md`
3. Bump the updated version and date of the release for manual `etime.1`
4. Commit these changes to a branch called by the version name – e.g. `v1.2.3`
5. Open then merge a pull request with these changes
6. Draft a [new release][releases] using the version name and entries from the
   `CHANGELOG.md`
7. Publish this as the latest release!
8. Close the current milestone for the latest release then create a new one

In deciding a version number, best judgement should be used to follow
[semantic versioning][semver].

### Signing notarizations

Packaging for the release process begins after a new version tag is created.

Builds for various targets are made with [goreleaser][goreleaser] then signed by
[gon][gon] and uploaded to the action artifacts.

Only compilations for macOS are signed at this time. Verifying binaries made for
other operating systems is left as an exercise for the developer.

#### Keychaining certificates

Certain credentials and certificates are requested for the signing processes.

Apple holds the keys for [developer credentials][credentials] and
[system certificates][certificates]. A "Developer ID Application" is needed on
the system keychain and any missing but matching certificates too.

Account information is also needed as environment variables in the `.env` file.

#### Processing packages

Signing and notarizing binaries is an automatic process that happens after
making a release build.

Special tooling and a macOS system is required for this process. Tooling can be
setup with a packaging flake:

```sh
$ flake develop .#gon
```

With the above ready the following commands will hopefully officiate things:

```sh
$ make release  # Build and notarize a release
$ gon .gon.hcl  # Troubleshoot specific errors
```

#### Verifying a signature

Unpackage the output disk image to make sure everything was successful with:

```sh
$ spctl -a -vvv -t install ./etime
```

[certificates]: https://www.apple.com/certificateauthority/
[credentials]: https://developer.apple.com/account/resources/certificates/list
[golang]: https://go.dev/dl/
[gon]: https://github.com/Bearer/gon
[goreleaser]: https://github.com/goreleaser/goreleaser
[learn_go]: https://go.dev/learn/
[nix]: https://zero-to-nix.com
[releases]: https://github.com/zimeg/emporia-time/releases
[runner]: https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners/adding-self-hosted-runners
[semver]: https://semver.org/spec/v2.0.0.html
[wiki]: https://github.com/zimeg/emporia-time/wiki
