package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getHome () string {
  homeDir, err := os.UserHomeDir() 
  if err != nil {
    fmt.Println("error getting user's home directory: ", err);
    return "Error getting user's home directory"
  }

  dirName := ".key-store"
  dirPath := filepath.Join(homeDir, dirName)

  return dirPath
}
