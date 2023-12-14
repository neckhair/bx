# bx

Unofficial command line client for [Bexio](https://bexio.com).

Features:

- List contacts

## Installation

Currently the only way to install `bx` is by `go install`:

```sh
go install github.com/neckhair/bx@main
```

It installs the tool into your `$GOPATH/bin`.

More installation options might follow later.

## Setup

Before you can run the CLI for the first time, you need to setup the credentials. You can create the
config file yourself, or you can just use the `setup` command.

This command will ask for the client id and the client secret, which you can get at
[Bexio's Developer Portal](https://developer.bexio.com).

As Redirect URL, use `http://localhost:50424`.

Then run the command:

```sh
bx setup
```
