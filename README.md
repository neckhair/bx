# bx

Unofficial command line client for [Bexio](https://bexio.com).

Features:

- List contacts

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
