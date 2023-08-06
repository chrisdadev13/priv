package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"github.com/urfave/cli/v2"
)

type Credentials struct {
  Name string `json: "name"`
  Pass string `json: "pass"`
  Date string `json: "date"`
}

func doWrite(c *cli.Context) error {
  var dir = c.String("dir")
  var name = c.String("name")
  var pass = c.String("pass")

  currentTime := time.Now()

  encryptedPassword, err := encrypt(pass)

  if err != nil {
    fmt.Println("Something wrong happened during encryption")
    return err
  }

  credentials := Credentials {
    Name: name,
    Pass: encryptedPassword,
    Date: currentTime.Format("2006-01-01"),
  }

  filePath := getHome() + "/" + dir + ".json"

  _, err = os.Stat(filePath); if err != nil {
    fmt.Printf("The store '%s' does not exist.\n", filePath)
    return err
  }

  file, err := os.Create(filePath)

  if err != nil {
    fmt.Println("There was an error writing in the file:", err)
    return err
  }

  defer file.Close()

  encoder := json.NewEncoder(file)
  err = encoder.Encode(credentials)

  if err != nil {
    fmt.Println("Error encoding JSON:", err)
  }

  fmt.Println("Credentials stored successfully in the store:", dir)

  return nil
}
