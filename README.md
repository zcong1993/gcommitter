# gcommitter

[![CircleCI](https://circleci.com/gh/zcong1993/gcommitter/tree/master.svg?style=svg)](https://circleci.com/gh/zcong1993/gcommitter/tree/master)
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
# build
$ go build ./...
# then move the output to your `$PATH` folder.
```

## License

MIT &copy; zcong1993
