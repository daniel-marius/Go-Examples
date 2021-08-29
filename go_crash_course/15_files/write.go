package main

import (
  "bufio"    
  "fmt"
  "log"
  "os"
  "io/ioutil"
)

func main() {
  write4()
}

// With Default Buffer Size of 4096 bytes
func write() {
  file, err := os.Create("./file.txt")
  if err != nil {
    log.Fatal(err)
  }

  writer := bufio.NewWriter(file)
  linesToWrite := []string{"This is an example1", "to show how", "to write to a file", "line by line."}
  for _, line := range linesToWrite {
    bytesWritten, err := writer.WriteString(line + "\n")
    if err != nil {
        log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
    }
    fmt.Printf("Bytes Written: %d\n", bytesWritten)
    fmt.Printf("Available: %d\n", writer.Available())
    fmt.Printf("Buffered : %d\n", writer.Buffered())
  }
  writer.Flush()
}

// With Custom Buffer Size of 10 bytes
func write2() {
  file, err := os.Create("./file.txt")
  if err != nil {
    log.Fatal(err)
  }

  writer := bufio.NewWriterSize(file, 10)
  linesToWrite := []string{"This is an example2", "to show how", "to write to a file", "line by line."}
  for _, line := range linesToWrite {
    bytesWritten, err := writer.WriteString(line + "\n")
    if err != nil {
        log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
    }
    fmt.Printf("Bytes Written: %d\n", bytesWritten)
    fmt.Printf("Available: %d\n", writer.Available())
    fmt.Printf("Buffered : %d\n", writer.Buffered())
  }
  writer.Flush()
}

func write3() {
  file, err := os.Create("./file.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  linesToWrite := []string{"This is an example3", "to show how", "to write to a file", "line by line."}
  for _, line := range linesToWrite {
    file.WriteString(line + "\n")
  }
}

func write4() {
  linesToWrite := "This is an example4 to show how to write to file using ioutil"
  err := ioutil.WriteFile("file.txt", []byte(linesToWrite), 0777)
  if err != nil {
    log.Fatal(err)
  }
}
