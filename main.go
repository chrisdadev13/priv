package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
	"github.com/chrisdadev13/priv/logger"
)

const version = "0.1.0"

var revision = "HEAD"

func main() {
	if err := newApp().Run(os.Args); err != nil {
		exitCode := 1
		if excoder, ok := err.(cli.ExitCoder); ok {
			exitCode = excoder.ExitCode()
		}
		logger.Log("error", err.Error())
		os.Exit(exitCode)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "priv"
	app.Usage = "Manage password and credentials locally"
	app.Version = fmt.Sprintf("%s (rev:%s)", version, revision)
	app.Authors = []*cli.Author{{
		Name:  "chrisdadev13",
		Email: "chrisdadev13@gmail.com",
	},}
  app.Commands = commands

  return app
}
