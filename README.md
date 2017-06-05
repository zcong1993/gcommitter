# gcommitter 

[![Wercker](https://img.shields.io/wercker/ci/wercker/docs.svg?style=flat-square)](https://app.wercker.com/project/byKey/a05e0a4c35ad0641d25d05bc685b2b2d)
[![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/gcommitter)](https://goreportcard.com/report/github.com/zcong1993/gcommitter)

> Easy way of git commit and push

just `git add -A && git commit -m "msg" && git push` in one command

## Usage:
> [Download](https://github.com/zcong1993/gcommitter/releases) the package, put in any `$PATH` folder.
```bash
# in your work folder
$ gcomitter [flag] [commit msg]
# example
$ gcomitter init # default msg is "backup"
# commit and push
$ gcomitter -p init
```

## Build:

```bash
$ git clone https://github.com/zcong1993/gcommitter.git
$ cd gcomitter
$ go install
$ go build gcomitter.go
# then move the output to your `$PATH` folder.
```

## License

MIT &copy; zcong1993
