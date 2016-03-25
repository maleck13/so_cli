package cmd

import (
	"github.com/maleck13/so_cli/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/maleck13/so_cli/cmd/subcmd"
)

//define our commands the logic is contained in the action function
func GetCmd() cli.Command {
	command := cli.Command{
		Name:  "get",
		Usage: "get a resource",
	}
	command.Subcommands = []cli.Command{
		subCmdUnanswered(),
		subCmdComments(),
	}

	return command
}

func subCmdUnanswered() cli.Command {

	return cli.Command{
		Name:        "unanswered",
		Description: "this command will get you all unanswered question tagged with --tags",
		Usage:       "get unanswered --tags=go;http",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "tags",
				Value:       "go",
				Usage:       "get the latest questions for <tag>. Passing more than one tag use tag1;tag2",
				Destination: &subcmd.FlagTag,
			},
		},
		Action: subcmd.GetUnanswered,
	}
}

func subCmdComments() cli.Command {
	return cli.Command{
		Name:        "comments",
		Description: "get all the comments from a question.",
		Usage:       "get comments <questionid>",
		Action:      subcmd.GetComments,
	}
}
