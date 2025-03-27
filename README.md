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
our [download site][downloads]. Besides that we also prepared repositories for
DEB and RPM packages which can be found at [Baltorepo][baltorepo]. If you prefer
to use containers you could use our images published on [GHCR][ghcr],
[Docker Hub][dockerhub] or [Quay][quay]. You are a Mac user? Just take a look
at our [homebrew formula][homebrew]. If you need further guidance how to
install this take a look at our [documentation][docs].

## Build

If you are not familiar with [Nix][nix] it is up to you to have a working
environment for Go (>= 1.24.0) and Nodejs (20.x) as the setup won't we covered
within this guide. Please follow the official install instructions for
[Go][golang] and [Nodejs][nodejs]. Beside that we are using [go-task][gotask] to
define all commands to build this project.

```console
git clone https://github.com/gopad/gopad-api.git
cd gopad-api

task fe:install fe:build be:build
./bin/gopad-api -h
```

If you got [Nix][nix] and [Direnv][direnv] configured you can simply execute
the following commands to get al dependencies including [go-task][gotask] and
the required runtimes installed. You are also able to directly use the process
manager of [devenv][devenv]:

```console
cat << EOF > .envrc
use flake . --impure --extra-experimental-features nix-command
EOF

direnv allow
```

We are embedding all the static assets into the binary so there is no need for
any webserver or anything else beside launching this binary.

## Development

To start developing on this project you have to execute only a few commands. To
start development just execute those commands in different terminals:

```console
task watch:server
task watch:frontend
```

The development server of the backend should be running on
[http://localhost:8080](http://localhost:8080) while the frontend should be
running on [http://localhost:5173](http://localhost:5173). Generally it supports
hot reloading which means the services are automatically restarted/reloaded on
code changes.

If you got [Nix][nix] configured you can simply execute the [devenv][devenv]
command to start the frontend, backend, MariaDB, PostgreSQL and Minio:

```console
devenv up
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
[downloads]: https://dl.gopad.eu
[baltorepo]: https://gopad.baltorepo.com/stable/
[homebrew]: https://github.com/gopad/homebrew-gopad
[ghcr]: https://github.com/orgs/gopad/packages
[dockerhub]: https://hub.docker.com/r/gopad/gopad-api/tags/
[quay]: https://quay.io/repository/gopad/gopad-api?tab=tags
[docs]: https://gopad.eu/
[nix]: https://nixos.org/
[golang]: http://golang.org/doc/install.html
[gotask]: https://taskfile.dev/installation/
[direnv]: https://direnv.net/
[devenv]: https://devenv.sh/
