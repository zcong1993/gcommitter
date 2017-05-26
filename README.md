# gcommitter

> Easy way of git commit and push

just `git add -A && git commit -m "msg" && git push` in one command

## Usage:
> [download](https://github.com/zcong1993/gcommitter/releases) the package, put in any `$PATH` folder.
```bash
# in your work folder
$ gcomitter [flag] [commit msg]
# example
$ gcomitter init
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
