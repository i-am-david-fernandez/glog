# glog

[![GitHub](https://img.shields.io/github/license/i-am-david-fernandez/glog.svg)](https://raw.githubusercontent.com/i-am-david-fernandez/glog/master/LICENSE)
[![Release](https://img.shields.io/github/release/i-am-david-fernandez/glog.svg?style=flat-square)](https://github.com/i-am-david-fernandez/glog/releases/latest)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/i-am-david-fernandez/glog)
[![Go Report Card](https://goreportcard.com/badge/github.com/i-am-david-fernandez/glog?style=flat-square)](https://goreportcard.com/report/github.com/i-am-david-fernandez/glog)

[![Build Status](https://travis-ci.com/i-am-david-fernandez/glog.svg?branch=master)](https://travis-ci.com/i-am-david-fernandez/glog)
[![GolangCI](https://golangci.com/badges/github.com/i-am-david-fernandez/glog.svg)](https://golangci.com)

Package `glog` is a go logging library.

Primarily, it is a convenience wrapper around `go-logging` (specifically the fork [shenwei356/go-logging](https://github.com/shenwei356/go-logging) that adds colour support for Windows).

`glog` provides convenience functions for configuring commonly-used logging backends (console and file) and for submitting log messages via a (package-scoped) global logger, akin to the print-style helper methods in the standard library log package.

It also includes additional backends: a convenience file-based backend and an (unlimited-size) in-memory list backend. This list backend is intended for use in relatively short-lived scenarios, such as batch-processing operations where the log output from each batch is to be treated independently (e.g., conditionally stored or transmitted). In such scenarios, one would clear the backend at the beginning of each batch run and decide what to do with the results at the end. A summary (the number of logged messages of each log level) is available to aid in conditional use.
