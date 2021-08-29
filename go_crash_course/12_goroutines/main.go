package main

import (
  "fmt"
  "runtime"
  "sync"
  // "time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
  // var msg = "Hello"
  // wg.Add(1)
  // go func(msg string) {
  //   fmt.Println(msg)
  //   wg.Done()
  // }(msg)
  // msg = "Goodbye"
  // time.Sleep(100 * time.Millisecond)
  // wg.Wait()

  runtime.GOMAXPROCS(100)
  for i := 0; i < 10; i++ {
    wg.Add(2)
    m.RLock()
    go sayHello()
    m.Lock()
    go increment()
  }
  wg.Wait()

  program1()
}

func sayHello() {
  fmt.Printf("Hello #%v\n", counter)
  m.RUnlock()
  wg.Done()
}

func increment() {
  counter++
  m.Unlock()
  wg.Done()
}
