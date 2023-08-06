package main 

import (
  "github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
  commandInit,
  commandCreate,
  commandWrite,
  commandGet,
}

var commandInit = &cli.Command{
  Name: "init",
  Usage: "Create a remote directory to save your stores",
  Action: doInit,
}

var commandCreate = &cli.Command{
  Name: "create",
  Usage: "Create a file to store your credentials",
  Action: doCreate,
}

var commandWrite = &cli.Command{
  Name: "write",
  Usage: "Store your credentials in one of your stores",
  Action: doWrite,
  Flags: []cli.Flag{
    &cli.StringFlag{
      Name: "dir",
      Aliases: []string{"d"},
      Usage: "Define the store where you want to save your credentials",
      Required: true,
    },
    &cli.StringFlag{
      Name: "name",
      Aliases: []string{"n"},
      Usage: "Define the username/email/id required to log in",
      Required: true,
    },
    &cli.StringFlag{
      Name: "pass",
      Aliases: []string{"p"},
      Usage: "Define the password",
      Required: true,
    },
  },
}

var commandGet = &cli.Command{
  Name: "get",
  Usage: "Get your password",
  Action: doGet,
  Flags: []cli.Flag{
    &cli.StringFlag{
      Name: "dir",
      Aliases: []string{"d"},
      Usage: "Define the store where you retrieve your credentials",
      Required: true,
    },
  },
}
