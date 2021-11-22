![GitHub](https://img.shields.io/github/license/mmmommm/goinit)
![GitHub CI Status](https://img.shields.io/github/workflow/status/mmmommm/goinit/ci?label=CI)
![GitHub Release Status](https://img.shields.io/github/workflow/status/mmmommm/goinit/Release?label=release)

# goinit
Generate initial configuration files for Go.

Here are generated by this cli tool.
```
- main.go
- README.md
- LICENSE
- .github/workflows/lint.yml
- .github/workflows/test.yml
- .gitignore
- .golangci.yml
```

## Required
- Go 1.16~

## Installation
```
$ go install github.com/mmmommm/goinit@latest
```

### MacOS
If you want to install on MacOS, you can use Homebrew.
```
brew install mmmommm/tap/goinit
```

## Usage
```sh
$ goinit
```

## Option
```sh
$ goinit p ${package_name}
```

Run `go mod init ${package_name}`


#### Here is not implemented now.
```sh
$ goinit d
```

Create dockerfile for Go.

## After run goinit

```
go run main.go
curl localhost:8080/Go
```

And then response this sentence.
>Hi there, I love Go

open LICENSE file and rewrite ${your_account_name}

```
Copyright (c) 2021 ${your_account_name}`
```
