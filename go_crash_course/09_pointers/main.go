package main

import (
  "fmt"
)

func main() {
  // var a int = 42
  // var b *int = &a
  // display the address where a is stored
  // fmt.Println(a, b)
  // a = 27
  // fmt.Println(a, b)
  // derefencing the pointer, display the value of b
  // fmt.Println(a, *b)

  // a := [3]int{1, 2, 3}
  // b := &a[0]
  // c := &a[1]
  // fmt.Printf("%v %p %p\n", a, *b, *c)

  var ms *myStruct
  // fmt.Println(ms)
  // ms = &myStruct{foo: 42}
  ms = new(myStruct)
  // (*ms).foo = 42
  // fmt.Println((*ms).foo)
  ms.foo = 42
  fmt.Println(ms.foo)

  // a := [3]int{1, 2, 3}
  // b := a
  // fmt.Println(a, b)
  // a[1] = 42
  // fmt.Println(a, b)

  a := map[string]string {"foo": "bar", "baz": "buz"}
  b := a
  fmt.Println(a, b)
  a["foo"] = "qux"
  fmt.Println(a, b)
}

type myStruct struct {
  foo int
}
