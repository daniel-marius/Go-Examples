package main

import (
  "fmt"
  "sync"
)

var wg = sync.WaitGroup{}

func main() {
  ch := make(chan int, 50)
  // for j := 0; j < 5; j++ {
  //   wg.Add(2)
  //
  //   go func() {
  //     i := <- ch
  //     fmt.Println(i)
  //     wg.Done()
  //   }()
  //
  //   go func() {
  //     ch <- 42
  //     wg.Done()
  //   }()
  // }

  wg.Add(2)

  // Receive go routine
  go func(ch <- chan int) {
    // i := <- ch
    // fmt.Println(i)
    // i = <- ch
    // fmt.Println(i)

    // for i := range ch {
    //   fmt.Println(i)
    // }

    for {
      if i, ok := <- ch; ok {
        fmt.Println(i)
      } else {
        break
      }
    }

    wg.Done()
  }(ch)

  // Send go routine
  go func(ch chan <- int) {
    ch <- 42
    ch <- 27
    // fmt.Println(<-ch)
    close(ch)
    wg.Done()
  }(ch)

  wg.Wait()
}
