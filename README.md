# gcommitter [![wercker status](https://app.wercker.com/status/a05e0a4c35ad0641d25d05bc685b2b2d/s/master "wercker status")](https://app.wercker.com/project/byKey/a05e0a4c35ad0641d25d05bc685b2b2d)

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
