
# tiant-micro

<p align="center">
<img align="center" width="120px" src="./doc/logo.png">
</p>

A lightweight micro service framework modeled after [go-zero](https://go-zero.dev/).

Based on [gRPC](https://grpc.io/) and [gin](https://gin-gonic.com/).

## Requirements
- Golang >= 1.20

## Use

1. Clone

```bash
git clone --depth 1 --branch main https://github.com/cbpantianya/tiant-micro.git <project-name>
cd <project-name>
go mod tidy
```

2. Rename

Rename all `config.ex.toml` to `config.toml`.

## Run 

### Docker
Make sure you are in the root path.

example:

```bash
docker build -t user:latest -f ./app/rpc/user/dockerfile .
```
There is a dockerfile template in `app/rpc/user/dockerfile`.



