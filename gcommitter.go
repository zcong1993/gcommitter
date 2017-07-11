package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"os"
)

// checkOut is a helper function check if out is not nil
func checkOut(out []byte) {
	if string(out) != "" {
		log.Fatal(string(out))
	}
}

// checkErr is a helper function panic err if err is not nil
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// excmd is a helper function execute cmd and return a combined output
func excmd(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func main() {
	var help, version, push bool
	var tag string
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&version, "v", false, "show version")
	flag.BoolVar(&push, "p", false, "commit and push")
	flag.StringVar(&tag, "t", "", "add and push tag")
	flag.Parse()
	if help {
		showHelp()
		os.Exit(0)
	}
	if version {
		fmt.Println(Version())
		os.Exit(0)
	}
	msg := strings.Join(flag.Args(), " ")
	if msg == "" {
		msg = "backup"
	}
	if tag != "" {
		out, err := excmd("git", "tag", "-a", tag, "-m", msg)
		checkErr(err)
		checkOut(out)
		out, err = excmd("git", "push", "origin", tag)
		checkErr(err)
		checkOut(out)
		if push {
			out, err = excmd("git", "push")
			checkErr(err)
			fmt.Printf("%s\n", out)
		}
		os.Exit(0)
	}
	out, err := excmd("git", "status", "--porcelain")
	checkErr(err)
	if string(out) == "" {
		log.Fatal("nothing to commit, working tree clean")
	}
	out, err = excmd("git", "add", "-A")
	checkErr(err)
	checkOut(out)
	out, err = excmd("git", "commit", "-m", msg, "--quiet")
	checkErr(err)
	checkOut(out)
	if push {
		out, err = excmd("git", "push")
		checkErr(err)
		fmt.Printf("%s\n", out)
	}
	fmt.Println("all done!")
}

func showHelp() {
	helpText := `
Usage :
	gct [flag] [commit msg]

Options:
	-p, --p 		commit and push
	-t=tag				run as this 'git tag -a [tag] -m [msg] && git push origin [tag]'
`
	fmt.Println(helpText)
}
