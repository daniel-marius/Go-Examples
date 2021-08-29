package main

import (
  "fmt"
  // "math"
)

func main() {
  // if true {
  //   fmt.Println("The test is true")
  // }

  // statePopulations := map[string]int{
  //   "City1": 43532423,
  //   "City2": 476574564,
  //   "City3": 46765656,
  // }
  //
  // if pop, ok := statePopulations["City3"]; ok {
  //   fmt.Println(pop)
  // }

  // number := 50
  // guess := 30
  //
  // if guess < 1 || returnTrue() || guess > 100 {
  //   fmt.Println("The guess must be between 1 and 100")
  // }
  //
  // if guess >= 1 && guess <= 100 {
  //   if guess < number {
  //     fmt.Println("Too low")
  //   }
  //
  //   if guess > number {
  //     fmt.Println("Too high")
  //   }
  //
  //   if guess == number {
  //     fmt.Println("Too ok")
  //   }
  //
  //   fmt.Println(number <= guess, number >= guess, number != guess)
  // }
  //
  // fmt.Println(!true)

  // myNum := 0.123
  // if myNum == math.Pow(math.Sqrt(myNum), 2) {
  //   fmt.Println("These are the same")
  // } else {
  //   fmt.Println("These are different")
  // }
  //
  // if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < 0.001 {
  //   fmt.Println("These are the same")
  // } else {
  //   fmt.Println("These are different")
  // }

  // switch 5 {
  // case 1, 5, 10:
  //   fmt.Println("one")
  // case 2, 4 ,6:
  //   fmt.Println("two")
  // default:
  //   fmt.Println("Not one or two")
  // }

  // switch i := 2 + 3;i {
  // case 1, 5, 10:
  //   fmt.Println("one")
  // case 2, 4 ,6:
  //   fmt.Println("two")
  // default:
  //   fmt.Println("Not one or two")
  // }

  // i := 10
  // switch {
  //   case i <= 10:
  //     fmt.Println("ten")
  //     // fallthrough
  //   case i >= 20:
  //     fmt.Println("twenty")
  //   default:
  //     fmt.Println("above twenty")
  // }

  // var i interface{} = 1
  // var i interface{} = "1"
  var i interface{} = [3]int{}
  switch i.(type) {
    case int:
      fmt.Println("i is an int")
      break
    case string:
      fmt.Println("i is a string")
    case [3]int:
      fmt.Println("i is an array of size 3")
    default:
      fmt.Println("i is another type")
  }

}

func returnTrue() bool {
  fmt.Println("Returning true")
  return true
}
