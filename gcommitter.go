package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func checkOut(out []byte) {
	if string(out) != "" {
		log.Fatal(out)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func excmd(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	return out, nil
}

func main() {
	help := flag.Bool("h", false, "show help")
	push := flag.Bool("p", false, "commit and push")
	flag.Parse()
	if *help {
		fmt.Println("\nUsage :\tgcommitter [flag] [commit msg]")
		fmt.Println("\nFlag :\t-p, --p, \tcommit and push")
		return
	}
	msg := strings.Join(flag.Args(), " ")
	if msg == "" {
		msg = "backup"
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
	if *push {
		out, err = excmd("git", "push")
		checkErr(err)
		fmt.Println(out)
	}
	fmt.Println("all done!")
}
