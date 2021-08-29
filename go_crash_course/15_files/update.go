package main

import (
  "fmt"
  "os"
  "time"
)

func main() {
  update()
}

func update() {
  fileName := "file.txt"
  currentTime := time.Now().Local()

  // Set both access time and modified time of the file to the current time
  err := os.Chtimes(fileName, currentTime, currentTime)
  if err != nil {
      fmt.Println(err)
  }
}
