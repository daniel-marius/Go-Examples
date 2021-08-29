package main

import (
  "fmt"
)

func main() {
  // for i := 0; i < 5; i++ {
  //   sayMessage("Hello Go!", i)
  // }

  // greeting := "hello"
  // name := "hello2"
  // sayGreeting(&greeting, &name)
  // fmt.Println(name)

  // s := sum(1, 2, 3, 4, 5)
  // fmt.Println("The sum is", *s)

  // d, err := divide(5.0, 3.0)
  // if err != nil {
  //   fmt.Println(err)
  //   return
  // }
  // fmt.Println(d)

  // var f func() = func() {
  //   fmt.Println("Hello Go!")
  // }
  // f()

  g := greeter {
    greeting: "Hello",
    name: "Go",
  }
  g.greet()
  fmt.Println("The new name is:", g.name)
}

func sayMessage(msg string, idx int) {
  fmt.Println(msg)
  fmt.Println("The value of the index is:", idx)
}

func sayGreeting(greeting, name *string) {
  fmt.Println(*greeting, *name)
  *name = "hello3"
  fmt.Println(*name)
}

func sum(values ...int) *int {
  fmt.Println(values)
  result := 0
  for _, v := range values {
    result += v
  }
  return &result
  // fmt.Println(msg, result)
}

func divide(a, b float64) (float64, error) {
  if b == 0.0 {
    return 0.0, fmt.Errorf("Cannot divide by zero")
  }
  return a / b, nil
}

type greeter struct {
  greeting string
  name string
}

func (g *greeter) greet() {
  fmt.Println(g.greeting, g.name)
  g.name = ""
}
