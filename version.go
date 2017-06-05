package main

import "fmt"

const (
	AppName = "gcommitter"
	AppVersion = "0.1.0"
)

func Version() string {
	return fmt.Sprintf("%s %s.\nCopyright (c) 2017, zcong1993.", AppName, AppVersion)
}

func Test() {
	fmt.Println("xsxs")
}
