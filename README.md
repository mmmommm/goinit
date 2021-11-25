![GitHub](https://img.shields.io/github/license/mmmommm/goinit)
![GitHub CI Status](https://img.shields.io/github/workflow/status/mmmommm/goinit/ci?label=CI)

# goinit
Generate initial configuration files for Go.

<img width="524" alt="example usage goinit command" src="https://user-images.githubusercontent.com/51479834/143152499-3e4dbd69-ded8-4121-8d58-a57f623bb4e0.png">

Here are generated by this cli tool.

- [main.go](https://github.com/mmmommm/goinit/blob/main/cmd/files/main.go)
- [README.md](https://github.com/mmmommm/goinit/blob/main/cmd/files/README.md)
- [LICENSE](https://github.com/mmmommm/goinit/blob/main/cmd/files/LICENSE)
- [.github/workflows/lint.yml](https://github.com/mmmommm/goinit/blob/main/cmd/files/lint.yml)
- [.github/workflows/test.yml](https://github.com/mmmommm/goinit/blob/main/cmd/files/test.yml)
- [.gitignore](https://github.com/mmmommm/goinit/blob/main/cmd/files/.gitignore)
- [.golangci.yml](https://github.com/mmmommm/goinit/blob/main/cmd/files/.golangci.yml)

## Required
- Go 1.16~

# Installation
```
$ go install github.com/mmmommm/goinit@latest
```

## MacOS
If you want to install on MacOS, you can use Homebrew.
```
brew install mmmommm/tap/goinit
```

#### upgrade
```
brew upgrade mmmommm/tap/goinit
```

## Windows, Linux etc...
Download the binary from [here](https://github.com/mmmommm/goinit/releases/tag/v0.1.4).

## Usage
```sh
$ goinit ${directory_name}

$ goinit example
```

## Option
```sh
$ goinit ${directory_name} -m ${package_name}

$ goinit example -m github.com/mmmommm/example
```
this option run `go mod init github.com/mmmommm/example`

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
