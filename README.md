# Gopad: API server

[![Build Status](https://cloud.drone.io/api/badges/gopad/gopad-api/status.svg)](https://cloud.drone.io/gopad/gopad-api)
[![Join the Matrix chat at https://matrix.to/#/#gopad:matrix.org](https://img.shields.io/badge/matrix-%23gopad-7bc9a4.svg)](https://matrix.to/#/#gopad:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/8592cd6c200d4e0cb2564c82498aaee1)](https://www.codacy.com/app/gopad/gopad-api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gopad/gopad-api&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/gopad/gopad-api?status.svg)](http://godoc.org/github.com/gopad/gopad-api)
[![Go Report](https://goreportcard.com/badge/github.com/gopad/gopad-api)](https://goreportcard.com/report/github.com/gopad/gopad-api)
[![](https://images.microbadger.com/badges/image/gopad/gopad-api.svg)](http://microbadger.com/images/gopad/gopad-api "Get your own image badge on microbadger.com")

**This project is under heavy development, it's not in a working state yet!**

Gopoad will be a simple web interface to write and update markdown-based documents. You can compare it with an Etherpad, just focused on markdown writing and formatting. I thought it's time to implement a shiny application with Go for the API and with VueJS for the UI.


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.gopad.tech/api). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/gopad/homebrew-gopad).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.11.

```bash
git clone https://github.com/gopad/gopad-api.git
cd gopad-api

make generate build

./bin/gopad-api -h
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
