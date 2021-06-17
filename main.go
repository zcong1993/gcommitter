package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
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
	fmt.Printf("%s %s\n%s\n", green("INFO"), greenNoBold(infoMsg), strings.Join(msg, "\n"))
}

func showErr(errMsg string, msg ...string) {
	red := color.New(color.FgRed, color.Bold).SprintfFunc()
	fmt.Printf("%s %s\n%s\n", red("ERROR"), errMsg, strings.Join(msg, "\n"))
}

// checkErr is a helper function panic err if err is not nil.
func checkErr(out []byte, err error) {
	if err != nil {
		showErr(err.Error(), string(out))
		os.Exit(1)
	}
}

// excmd is a helper function execute cmd and return a combined output.
func excmd(name string, arg ...string) ([]byte, error) {
	info(fmt.Sprintf("%s %s", name, strings.Join(arg, " ")))
	cmd := exec.Command(name, arg...)
	return cmd.CombinedOutput()
}

func expectEmpty(out []byte, err error) {
	checkErr(out, err)

	if str := string(out); str != "" {
		showErr("output un empty", str)
	}
}

func showOut(out []byte, err error) {
	checkErr(out, err)
	info("output", string(out))
}

func main() {
	app := &cli.App{
		Name:        "gcommitter",
		UsageText:   "gcommitter [options] [commit messages...]",
		Description: "Git add + commit + push",
		Version:     buildVersion(version, commit, date, builtBy),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "push",
				Aliases: []string{"p"},
				Usage:   "if push",
			},
			&cli.StringFlag{
				Name:    "tag",
				Aliases: []string{"t"},
				Usage:   "add tag",
			},
		},
		Action: func(c *cli.Context) error {
			msg := strings.Join(c.Args().Slice(), " ")
			return process(msg, c.String("tag"), c.Bool("push"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		showErr(err.Error())
		os.Exit(1)
	}
}

func process(msg, tag string, push bool) error {
	if tag != "" {
		args := []string{"tag"}

		if msg != "" {
			args = append(args, "-a", tag, "-m", msg)
		} else {
			args = append(args, tag)
		}
		expectEmpty(excmd("git", args...))
		showOut(excmd("git", "push", "origin", tag))
		if push {
			showOut(excmd("git", "push"))
		}
		return nil
	}

	if msg == "" {
		return errors.New("commit message is required")
	}

	out, err := excmd("git", "status", "--porcelain")
	checkErr(out, err)
	if string(out) == "" {
		return errors.New("nothing to commit, working tree clean")
	}
	expectEmpty(excmd("git", "add", "-A"))
	showOut(excmd("git", "commit", "--quiet", "-m", msg))
	if push {
		showOut(excmd("git", "push"))
	}
	info("all done!")
	return nil
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
