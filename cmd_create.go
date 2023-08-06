package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
)

func doCreate(c *cli.Context) error{
  var name = c.Args().Get(0)
  dirPath := getHome()

  if name != ""{
    createStore(dirPath, name)
    return nil
  }

  fmt.Println("you need to provide a name for the store")
  return nil
}

func createStore(dirPath string, fileName string) error {
  filePath := dirPath + "/" + fileName + ".json"

  _, err := os.Stat(filePath); if err == nil {
		fmt.Printf("File '%s' already exists.\n", filePath)
		fmt.Printf("Try again with a different file name\n")
    return err
	} 

  file, err := os.Create(filePath)

  if err != nil { 
    fmt.Println("error creating file")
    return err 
  }
  defer file.Close()

  fmt.Println("file created successfully:", filePath)

  return nil
}
