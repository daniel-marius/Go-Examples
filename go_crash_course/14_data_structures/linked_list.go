package main

import "fmt"

func main() {
  singleList := initList()
  fmt.Printf("AddFront: A\n")
  singleList.AddFront("A")
  fmt.Printf("AddFront: B\n")
  singleList.AddFront("B")
  fmt.Printf("AddBack: C\n")
  singleList.AddBack("C")

  fmt.Printf("Size: %d\n", singleList.Size())

  err := singleList.Traverse()
  if err != nil {
     fmt.Println(err.Error())
  }

  fmt.Printf("RemoveFront\n")
  err = singleList.RemoveFront()
  if err != nil {
     fmt.Printf("RemoveFront Error: %s\n", err.Error())
  }

  fmt.Printf("RemoveBack\n")
  err = singleList.RemoveBack()
  if err != nil {
     fmt.Printf("RemoveBack Error: %s\n", err.Error())
  }

  fmt.Printf("RemoveBack\n")
  err = singleList.RemoveBack()
  if err != nil {
     fmt.Printf("RemoveBack Error: %s\n", err.Error())
  }

  fmt.Printf("RemoveBack\n")
  err = singleList.RemoveBack()
  if err != nil {
     fmt.Printf("RemoveBack Error: %s\n", err.Error())
  }

  err = singleList.Traverse()
  if err != nil {
     fmt.Println(err.Error())
  }

  fmt.Printf("Size: %d\n", singleList.Size())
}

type el struct {
  name string
  next *el
}

type singleList struct {
  len int
  head *el
}

func initList() *singleList {
  return &singleList {}
}

func (s *singleList) AddFront(name string) {
  el := &el {
    name: name,
  }

  if s.head == nil {
    s.head = el
  } else {
    el.next = s.head
    s.head = el
  }

  s.len += 1
  return
}

func (s *singleList) AddBack(name string) {
  el := &el {
    name: name,
  }

  if s.head == nil {
    s.head = el
  } else {
    current := s.head
    for current.next != nil {
      current = current.next
    }

    current.next = el
  }

  s.len += 1
  return
}

func (s *singleList) RemoveFront() error {
  if s.head == nil {
    return fmt.Errorf("List is empty")
  }

  s.head = s.head.next
  s.len -= 1
  return nil
}

func (s *singleList) RemoveBack() error {
  if s.head == nil {
    return fmt.Errorf("Remove back: List is empty")
  }

  var prev *el
  current := s.head
  for current.next != nil {
    prev = current
    current = current.next
  }

  if prev != nil {
    prev.next = nil
  } else {
    s.head = nil
  }

  s.len -= 1
  return nil
}

func (s *singleList) Front() (string, error) {
  if s.head == nil {
    return "", fmt.Errorf("Single list is empty")
  }

  return s.head.name, nil
}

func (s *singleList) Size() int {
  return s.len
}

func (s *singleList) Traverse() error {
  if s.head == nil {
    return fmt.Errorf("Traverse error: List is empty")
  }

  current := s.head
  for current != nil {
    fmt.Println(current.name)
    current = current.next
  }

  return nil
}
