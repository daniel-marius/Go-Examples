package main

import (
  "fmt"
  //"io/ioutil"
  "log"
  //"net/http"
)

func main() {
  // defer fmt.Println("aaa")
  // defer fmt.Println("bbb")
  // defer fmt.Println("ccc")

  // res, err := http.Get("http://www.google.com/robots.txt")
  // if err != nil {
  //   log.Fatal(err)
  // }
  //
  // robots, err := ioutil.ReadAll(res.Body)
  // defer res.Body.Close()
  // if err != nil {
  //   log.Fatal(err)
  // }
  // fmt.Printf("%s", robots)

  // a := "start"
  // defer fmt.Println(a)
  // a = "end"

  // a, b := 1, 0
  // ans := a / b
  // fmt.Println(ans)

  fmt.Println("start")
  defer func() {
    if err := recover(); err != nil {
      log.Println("Error: ", err)
    }
  }()
  panic("Something went wrong!")
  fmt.Println("end")

  // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //   w.Write([]byte("Hello Go!"))
  // })
  // err := http.ListenAndServe(":8080", nil)
  // if err != nil {
  //   panic(err.Error())
  // }
}
