# goBrute - Password Brute Forcer

[![Go Report Card](https://goreportcard.com/badge/github.com/tomp332/gobrute)](https://goreportcard.com/report/github.com/tomp332/gobrute)
[![Build](https://github.com/tomp332/gobrute/actions/workflows/main.yml/badge.svg)](https://github.com/tomp332/gobrute/actions/workflows/main.yml)
[![License](https://img.shields.io/github/license/tomp332/gobrute.svg)](https://github.com/tomp332/gobrute/blob/main/LICENSE.md)
[![Code Coverage](https://codecov.io/gh/tomp332/gobrute/settings/badge.svg)](https://codecov.io/gh/tomp332/gobrute?branch=main)

<div align="center"> 
  <img src="https://github.com/tomp332/gobrute/assets/47506972/893e19a0-9ee8-47a7-93e9-84fa617eb163" width="500" height="450" />
</div>

## Description

GoLang Brute Forcer is a command-line tool that performs brute force attacks to test the security of passwords or keys.
It is designed to help identify weak passwords and improve overall system security.


## Features

- Customizable character sets and password length.
- Parallel processing for faster brute-forcing.
- Supports different hashing algorithms (MD5, SHA-256, etc.).
- Customizable output format.

## Installation

### Pre-built binaries

You can download pre-built binaries for your platform from
the [Releases](https://github.com/your-username/golang-bruteForcer/releases) page. Extract the archive and add the
binary to your PATH.

### Build from source

To build from source, make sure you have Go (1.15+) installed and run the following commands:

```bash
$ git clone https://github.com/your-username/golang-bruteForcer.git
$ cd golang-bruteForcer
$ go build
