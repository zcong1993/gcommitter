package main

import "fmt"

var (
	// AppName is the cli name
	AppName = "gcommitter"
	version = "0.1.0"
	commit string
)

// Version show the cli's current version
func Version() string {
	return fmt.Sprintf("\n%s %s(%s).\nCopyright (c) 2017, zcong1993.", AppName, version, commit)
}
