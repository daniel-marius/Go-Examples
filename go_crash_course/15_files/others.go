package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

var (
  fileInfo *os.FileInfo
  err error
)

func main() {
  // currentDirectory, err := os.Getwd()
  // if err != nil {
  //     log.Fatal(err)
  // }
  // iterateFolder(currentDirectory)
  fileStats()
}

// Check if file exists
func checkFile() {
  info, err := os.Stat("temp")
  if os.IsNotExist(err) {
      log.Fatal("File does not exist.")
  }
  if info.IsDir() {
      fmt.Println("temp is a directory")
  } else {
      fmt.Println("temp is a file")
  }
}

// Check if file is empty
func checkFile2() {
  file, err := os.Create("emptyFile.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
}

// Check if folder exists
func checkFolder() {
  folderInfo, err := os.Stat("temp")
  if os.IsNotExist(err) {
      log.Fatal("Folder does not exist.")
  }
  log.Println(folderInfo)
}

func iterateFolder(path string) {
  filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("File Name: %s\n", info.Name())
    return nil
  })
}

func fileStats() {
  // Create a file
  file, err := os.Create("temp.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  // Write something to the file
  file.WriteString("some sample text" + "\n")

  // Gets stats of the file
  stats, err := os.Stat("temp.txt")
  if err != nil {
      log.Fatal(err)
  }

  // Prints stats of the file
  fmt.Printf("Permission: %s\n", stats.Mode())
  fmt.Printf("Name: %s\n", stats.Name())
  fmt.Printf("Size: %d\n", stats.Size())
  fmt.Printf("Modification Time: %s\n", stats.ModTime())
}
