# lettergen 
[![GitHub release](http://img.shields.io/github/release/jakewarren/lettergen.svg?style=flat-square)](https://github.com/jakewarren/lettergen/releases])
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/lettergen/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/lettergen)](https://goreportcard.com/report/github.com/jakewarren/lettergen)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)
> generates a letter avatar

A command line utility for [github.com/disintegration/letteravatar](https://github.com/disintegration/letteravatar)

## Install
### Option 1: Binary

Download the latest release from [https://github.com/jakewarren/lettergen/releases/latest](https://github.com/jakewarren/lettergen/releases/latest)

### Option 2: From source

```
go get github.com/jakewarren/lettergen
```

## Usage

```
❯ lettergen -h
usage: lettergen [<flags>] <letter>

cmd line tool to create single-letter avatar images.

Optional flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
  -o, --output-file=OUTPUT-FILE  The name of the image file.
  -s, --size=75                  the size in pixels for the image.
  -V, --version                  Show application version.

Args:
  <letter>  the letter to put in the logo.
```

## Changes

All notable changes to this project will be documented in the [changelog].

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this project adheres to [Semantic Versioning](http://semver.org/).

## License

MIT © 2018 Jake Warren

[changelog]: https://github.com/jakewarren/lettergen/blob/master/CHANGELOG.md
