# Gopad: API server

[![General Workflow](https://github.com/gopad/gopad-api/actions/workflows/general.yml/badge.svg)](https://github.com/gopad/gopad-api/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#gopad:matrix.org](https://img.shields.io/badge/matrix-%23gopad-7bc9a4.svg)](https://matrix.to/#/#gopad:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/7143ea13bd644aa3be6749ca967be7d0)](https://app.codacy.com/gh/gopad/gopad-api/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/gopad/gopad-api.svg)](https://pkg.go.dev/github.com/gopad/gopad-api) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/gopad/gopad-api)

> [!CAUTION]
> This project is in active development and does not provide any stable release
> yet, you can expect breaking changes until our first real release!

Gopoad will be a simple web interface to write and update markdown-based
documents. You can compare it with an Etherpad, just focused on markdown writing
and formatting. I thought it's time to implement a shiny application with Go for
the API and with VueJS for the UI.

## Install

You can download prebuilt binaries from the [GitHub releases][releases] or from
our [download site][downloads]. If you prefer to use containers you could use
our images published on [Docker Hub][dockerhub] or [Quay][quay]. You are a Mac
user? Just take a look at our [homebrew formula][homebrew]. If you need further
guidance how to install this take a look at our [documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.24, at least that's the version we are using.

```console
git clone https://github.com/gopad/gopad-api.git
cd gopad-api

make generate build
./bin/gopad-api -h
```

## Security

If you find a security issue please contact
[gopad@webhippie.de](mailto:gopad@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[releases]: https://github.com/gopad/gopad-api/releases
[downloads]: https://dl.gopad.eu/api
[homebrew]: https://github.com/gopad/homebrew-gopad
[dockerhub]: https://hub.docker.com/r/gopad/gopad-api/tags/
[quay]: https://quay.io/repository/gopad/gopad-api?tab=tags
[docs]: https://gopad.eu/
[golang]: http://golang.org/doc/install.html
