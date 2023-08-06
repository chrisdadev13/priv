package main

import (
  "fmt"
  "os"
  "encoding/json"
	"github.com/urfave/cli/v2"
)

func doGet(c *cli.Context) error {
  var dir = c.String("dir")
  
  filePath := getHome() + "/" + dir + ".json"

  file, err := os.Open(filePath)
  if err != nil {
    fmt.Printf("There was an error opening '%s'", filePath)
    return err
  }

  defer file.Close()

  var credentials Credentials

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&credentials)

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}

  if err != nil {
    fmt.Println("Something wrong happened during decryption")
    return err
  }

  decryptedPassword, err := decrypt(credentials.Pass)

  fmt.Printf("Your password for '%s' is: '%s'", dir, decryptedPassword)

  return nil
}
