package clihandler

import (
	"github.com/urfave/cli/v2"
)

func New() *cli.App {
	cliApp := cli.NewApp()

	return cliApp
}
