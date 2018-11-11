# columbus-cli

## About columbus-cli

columbus-cli is a command line interface to interact with the [Columbus] monitoring app.

## Development

To start development on columbus-cli execute the following steps.

### Setup Go and Dep

``` shell
brew install go
```

``` shell
brew install dep
```

``` shell
export GOPATH=$HOME/go
```

### Get the columbus-cli github project

``` shell
go get -v github.com/n3xtdata/columbus-cli
```

``` shell
cd $HOME/go/src/github.com/n3xtdata/columbus-cli
```

### Build

``` shell
make build
```
The build files can be found at $HOME/go-builds/columbus-cli.

### Install

``` shell
make install
```

The binary will be installed at "/usr/local/bin".

## Use the columbus cli

``` shell
columbus help
```

[Columbus]: https://www.github.com/n3xtdata/columbus