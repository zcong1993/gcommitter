package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"

	"github.com/fatih/color"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

func info(infoMsg string, msg ...string) {
	green := color.New(color.FgGreen, color.Bold).SprintfFunc()
	greenNoBold := color.New(color.FgGreen).SprintfFunc()
	fmt.Printf("%s %s\n%s\n\n", green("INFO"), greenNoBold(infoMsg), strings.Join(msg, "\n"))
}

func showErr(errMsg string, msg ...string) {
	red := color.New(color.FgRed, color.Bold).SprintfFunc()
	fmt.Printf("%s %s\n%s\n\n", red("ERROR"), errMsg, strings.Join(msg, "\n"))
}

func isLintStageOut(out string) bool {
	return strings.Contains(out, "lint-staged")
}

// checkErr is a helper function panic err if err is not nil
func checkErr(err error, out []byte) {
	str := string(out)
	if err != nil {
		showErr(err.Error(), str)
		os.Exit(1)
	}
}

// excmd is a helper function execute cmd and return a combined output
func excmd(name string, arg ...string) ([]byte, error) {
	info(fmt.Sprintf("%s %s", name, strings.Join(arg, " ")))
	cmd := exec.Command(name, arg...)
	return cmd.CombinedOutput()
}

func expectEmpty(out []byte, err error) {
	checkErr(err, out)

	str := string(out)
	if str != "" {
		if isLintStageOut(str) {
			info("", str)
			return
		}
		showErr("output un empty", str)
	}
}

func showOut(out []byte, err error) {
	checkErr(err, out)
	info("", string(out))
}

func main() {
	var help, showVersion, push bool
	var tag string
	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&push, "p", false, "commit and push")
	flag.StringVar(&tag, "t", "", "add and push tag")
	flag.Parse()

	if help {
		showHelp()
		os.Exit(0)
	}
	if showVersion {
		fmt.Println(buildVersion(version, commit, date, builtBy))
		os.Exit(0)
	}
	msg := strings.Join(flag.Args(), " ")
	if msg == "" {
		msg = "backup"
	}
	if tag != "" {
		expectEmpty(excmd("git", "tag", "-a", tag, "-m", msg))
		showOut(excmd("git", "push", "origin", tag))
		if push {
			showOut(excmd("git", "push"))
		}
		os.Exit(0)
	}

	out, err := excmd("git", "status", "--porcelain")
	checkErr(err, out)
	if string(out) == "" {
		showErr("nothing to commit, working tree clean")
		os.Exit(1)
	}
	expectEmpty(excmd("git", "add", "-A"))
	expectEmpty(excmd("git", "commit", "-m", msg, "--quiet"))
	if push {
		showOut(excmd("git", "push"))
	}
	info("all done!")
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

func buildVersion(version, commit, date, builtBy string) string {
	result := version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
		result = fmt.Sprintf("%s\nmodule version: %s, checksum: %s", result, info.Main.Version, info.Main.Sum)
	}
	return result
}
