package main

import (
  "fmt"
  "strconv"
)

// var i int = 42

// var (
//   actorName string = "hjvhj";
//   doctorNumber int = 3
//   season int = 11
// )
//
// var (
//   counter int = 0
// )

// var I int = 48

// var i int = 27

func main() {
  // var i int
  // i = 42
  // var j int = 27
  // k := 90
  // fmt.Printf("%v, %T", i, i)

  // fmt.Println(i)
  // var i int = 42
  // fmt.Println(i)

  var i int = 42
  fmt.Printf("%v, %T\n", i, i)

  // var j float32
  // var j int
  var j string
  j = strconv.Itoa(i)
  fmt.Printf("%v, %T\n", j, j)
}
