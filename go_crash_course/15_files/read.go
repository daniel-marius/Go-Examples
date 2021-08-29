package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  // WordbyWordScan()
  LinebyLineScan()
}

func WordbyWordScan() {
  file, err := os.Open("./file.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func LinebyLineScan() {
  file, err := os.Open("./file.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      fmt.Println(scanner.Text())
  }
  if err := scanner.Err(); err != nil {
      log.Fatal(err)
  }
}
