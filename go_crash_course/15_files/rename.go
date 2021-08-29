package main

import (
  "log"
  "os"
)

func main() {
  renameFile()
}

func renameFile() {
  // Create a file
  file, err := os.Create("file.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  // Change permission so that it can be moved
  err = os.Chmod("file.txt", 0777)
  if err != nil {
      log.Println(err)
  }


  err = os.Rename("file.txt", "newFile.txt")
  if err != nil {
      log.Fatal(err)
  }
}

func renameFolder() {
  // Create a directory
  err := os.Mkdir("temp", 0755)
  if err != nil {
      log.Fatal(err)
  }
  err = os.Rename("temp", "newTemp")
  if err != nil {
      log.Fatal(err)
  }
}
