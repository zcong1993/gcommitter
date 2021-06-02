package main

import "fmt"

var (
	// AppName is the cli name
	AppName = "gcommitter"
	version = "1.0.2"
)

// Version show the cli's current version
func Version() string {
	return fmt.Sprintf("\n%s %s.\nCopyright (c) 2021, zcong1993.", AppName, version)
}
