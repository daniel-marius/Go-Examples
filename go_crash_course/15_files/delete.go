package main

import (
  "log"
  "os"
)

func main() {
  delete()
}

func delete() {
  // err := os.Remove("./file.txt") // delete file
  err := os.Remove("folder") // delete folder
  if err != nil {
      log.Fatal(err)
  }
}
