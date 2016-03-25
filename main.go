package main

import (
	"github.com/maleck13/so_cli/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/maleck13/so_cli/cmd"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Usage = "gets and prints stackoverflow questions"
	app.Commands = []cli.Command{
		cmd.GetCmd(),
		cmd.SearchCmd(),
	}
	app.Run(os.Args)
}
