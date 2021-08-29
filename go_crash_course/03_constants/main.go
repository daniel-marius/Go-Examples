package main

import (
  "fmt"
)

// const a int16 = 27 // declared at the package level

// const a = iota // special character, enumereted constant

const (
  _ = iota + 5
  a
  b // b related constant
  c // related
)

const (
  _ = iota // ignore first value by assigning to blank identifier
  KB = 1 << (10 * iota)
  MB
  GB
)

func main() {
  // // internal constant
  // const myConst int = 42
  //
  // // exported constant
  // const MY_CONST int = 78
  //
  // fmt.Printf("%v, %T\n", myConst, myConst)

  // const a int = 14
  // const a = 42
  // var b int16 = 27
  // fmt.Printf("%v, %T\n", a + b, a + b)

  var d int
  fmt.Printf("%v, %T\n", a, a)
  fmt.Printf("%v, %T\n", d == b, d)
  fmt.Printf("%v, %T\n", c, c)

  fileSize := 400000000.
  fmt.Printf("%2fGB", fileSize/GB)

  aa := 1 << 0

  fmt.Println( aa)
}
