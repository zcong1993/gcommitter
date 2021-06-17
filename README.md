# gcommitter
<!--
[![Go Report Card](https://goreportcard.com/badge/github.com/zcong1993/gcommitter)](https://goreportcard.com/report/github.com/zcong1993/gcommitter)
-->

> Git add + commit + push

## Install

```bash
brew install zcong1993/homebrew-tap/gcommitter
```

## Usage

```bash
gcommitter -h

gcommitter -p "commit message"
# git add . && git commit -m "commit message" && git push

gcommitter -p -t v2.0.0
# git tag v2.0.0 && git push origin v2.1.0 && git push
```

## License

MIT &copy; zcong1993
