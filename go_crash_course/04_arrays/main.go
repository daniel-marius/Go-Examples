package main

import (
  "fmt"
)

func main() {
  // grades := [...]int{ 97, 98, 99 }
  // var students [3]string
  // fmt.Printf("%v\n", students)
  // students[0] = "std1"
  // students[2] = "std3"
  // students[1] = "std2"
  // fmt.Printf("%v\n", students[1])
  // fmt.Printf("%v", len(students))

  // var identityMatrix [3][3]int
  // identityMatrix[0] = [3]int{1, 0, 0}
  // identityMatrix[1] = [3]int{0, 1, 0}
  // identityMatrix[2] = [3]int{0, 0, 1}
  // fmt.Println(identityMatrix)

  a := [...]int{1, 2, 3} // array
  // a := []{1, 2, 3} // slice
  b := a // literal copy, pointing to a different set of data
  // b := &a // ref copy, pointing to the same set of data
  a[1] = 90
  // b[1] = 5
  fmt.Println(a)
  fmt.Println(b)
  fmt.Printf("Length: %v\n", len(a))
  fmt.Printf("Capacity: %v\n", cap(a))

  // a := [...]int{1, 2, 4, 5, 6, 7, 8, 9, 10}
  // b := a[:] // slice of all elements
  // c := a[3:] // slice from 4th element to last
  // d := a[:6] // slice first 6 elements
  // e := a[3:6] // slice the 4th, 5th, and 6th elements
  // a[5] = 42 // each slice points to the same underlying array
  // fmt.Println(a)
  // fmt.Println(b)
  // fmt.Println(c)
  // fmt.Println(d)
  // fmt.Println(e)

  // a := make([]int, 3, 100) // type, length, capacity
  // a := []int{}
  // fmt.Println(a)
  // fmt.Printf("Length: %v\n", len(a))
  // fmt.Printf("Capacity: %v\n", cap(a))
  // a = append(a, 1)
  // fmt.Println(a)
  // fmt.Printf("Length: %v\n", len(a))
  // fmt.Printf("Capacity: %v\n", cap(a))
  // // a = append(a, 2, 3, 4, 5)
  // a = append(a, []int{2, 3, 4, 5}...)
  // fmt.Println(a)
  // fmt.Printf("Length: %v\n", len(a))
  // fmt.Printf("Capacity: %v\n", cap(a))

  // a := []int{1, 2, 3, 4, 5}
  // fmt.Println(a)
  // b := a[1:] // removes the first element
  // b := a[:len(a) - 1] // remove the last element
  // b := append(a[:2], a[3:]...) // removes element on index 2
  // fmt.Println(b)
  // fmt.Println(a)
}
