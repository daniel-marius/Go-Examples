package main

import (
  "fmt"
  "reflect"
)

// type Doctor struct {
//   number int
//   actorName string
//   episodes []string
//   companions []string
// }

type Animal struct {
  name string `required max:"100"`
  origin string
}

type Bird struct {
  Animal
  speed float32
  canFly bool
}

func main() {
  // statePopulations := make(map[string]int)
  // statePopulations = map[string]int{
  //   "City1": 43532423,
  //   "City2": 476574564,
  //   "City3": 46765656,
  // }
  // statePopulations := map[string]int{
  //   "City1": 43532423,
  //   "City2": 476574564,
  //   "City3": 46765656,
  // }

  // fmt.Println(statePopulations)
  // statePopulations["City4"] = 98666986
  // fmt.Println(statePopulations["City4"])
  // delete(statePopulations, "City4")

  // pop, ok := statePopulations["City3"];
  // fmt.Println(pop, ok)
  // sp := statePopulations
  // sp := &statePopulations
  // fmt.Println(statePopulations)
  // delete(sp, "City3")
  // fmt.Println(sp)
  // fmt.Println(statePopulations)
  // fmt.Println(len(statePopulations))

  // aDoctor := Doctor {
  //   number: 3,
  //   actorName: "name1",
  //   companions: []string {
  //     "comp1",
  //     "comp2",
  //   },
  //   episodes: []string {},
  // }
  //
  // fmt.Println(aDoctor)

  // aDoctor := struct{name string}{name: "name1"}
  // anotherDoctor := aDoctor
  // anotherDoctor.name = "name2"
  // fmt.Println(aDoctor)
  // fmt.Println(anotherDoctor)

  // b := Bird{}
  // b.name = "bird"
  // b.origin = "EU"
  // b.speed = 56
  // b.canFly = false
  b := Bird{
    Animal: Animal{name: "bird", origin: "US"},
    speed: 67,
    canFly: false,
  }
  t := reflect.TypeOf(b)
  field, _ := t.FieldByName("name")
  fmt.Println(field.Tag)
}
