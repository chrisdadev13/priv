package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
)

func doInit(c *cli.Context) error {
  dirPath := getHome()

  if _, err := os.Stat(dirPath); os.IsNotExist(err) {
    createDir(dirPath)
    return nil
  }

  fmt.Println("error, the main directory already exist.")

  return nil
}

func createDir(dirPath string) error{
  err := os.MkdirAll(dirPath, 0755)
  if err != nil {
    fmt.Println("error creating store directory: ", err)
    return err
  }

  fmt.Println("password store created successfully: ", dirPath)

  return nil
}

