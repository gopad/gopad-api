# Gopad: API server

[![Build Status](http://drone.gopad.tech/api/badges/gopad/gopad-api/status.svg)](http://drone.gopad.tech/gopad/gopad-api)
[![Stories in Ready](https://badge.waffle.io/gopad/gopad-api.svg?label=ready&title=Ready)](http://waffle.io/gopad/gopad-api)
[![Join the Matrix chat at https://matrix.to/#/#gopad:matrix.org](https://img.shields.io/badge/matrix-%23gopad%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#gopad:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/8592cd6c200d4e0cb2564c82498aaee1)](https://www.codacy.com/app/gopad/gopad-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gopad/gopad-api&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/gopad/gopad-api?status.svg)](http://godoc.org/github.com/gopad/gopad-api)
[![Go Report](https://goreportcard.com/badge/github.com/gopad/gopad-api)](https://goreportcard.com/report/github.com/gopad/gopad-api)
[![](https://images.microbadger.com/badges/image/gopad/gopad-api.svg)](http://microbadger.com/images/gopad/gopad-api "Get your own image badge on microbadger.com")


**This project is under heavy development, it's not in a working state yet!**

TBD


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.gopad.tech/api). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/gopad/homebrew-gopad).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). As this project relies on vendoring of the dependencies you have to use a Go version `>= 1.6`. It is also possible to just simply execute the `go get github.com/gopad/gopad-api/cmd/gopad-api` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/gopad/gopad-api
cd $GOPATH/src/github.com/gopad/gopad-api
make clean build

./gopad-api -h
```


## Security

If you find a security issue please contact gopad@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
