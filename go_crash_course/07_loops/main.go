package main

import (
  "fmt"
)

func main() {
  // for i := 0; i < 5; i++ {
  //   fmt.Println(i)
  // }
  //
  // for i, j := 0, 0; i < 5; i, j = i + 1, j + 2 {
  //   fmt.Println(i, j)
  // }

  // for i := 0; i < 5; i++ {
  //   fmt.Println(i)
  //   if i % 2 == 0 {
  //     i /= 2
  //   } else {
  //     i = 2*i + 1
  //   }
  // }

  // i := 0
  // for ; i < 5; {
  //   fmt.Println(i)
  //   i++
  //   if i == 4 {
  //     break
  //   }
  // }

  // for i := 0; i < 10; i++ {
  //   if i % 2 == 0 {
  //     continue
  //   }
  //   fmt.Println(i)
  // }

// Loop:
//     for i := 1; i <= 3; i++ {
//       for j := 1; j <= 3; j++ {
//         fmt.Println(i * j)
//         // break the inner loop
//         // if i * j >= 3 {
//         //   break
//         // }
//
//         // break the outer loop
//         if i * j >= 3 {
//           break Loop
//         }
//       }
//     }

  // s := []int{1, 2, 4}
  // for k, v := range s {
  //   fmt.Println(k, v)
  // }

  s := "Hello Go!"
  for k, v := range s {
    // key and unicode representation for each character
    fmt.Println(k, v)

    // key and character
    fmt.Println(k, string(v))
  }
}
