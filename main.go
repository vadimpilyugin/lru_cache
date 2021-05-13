package main

import "fmt"

func main() {
  c := NewCache(3) // {}
  c.Put(1, "str1") // {1: “str1”}
  fmt.Printf("%v\n", c)
  c.Put(2, "str2") // {1: “str1”, 2: “str2”}
  fmt.Printf("%v\n", c)
  c.Put(3, "str3") // {1: “str1”, 2: “str2”, 3: “str3”}
  fmt.Printf("%v\n", c)
  fmt.Println("c.Get(3) = ", c.Get(3))
  fmt.Printf("%v\n", c)
  fmt.Println("c.Get(2) = ", c.Get(2))
  fmt.Printf("%v\n", c)
  fmt.Println("c.Get(1) = ", c.Get(1))
  fmt.Printf("%v\n", c)
  fmt.Println("c.Get(3) = ", c.Get(3))
  fmt.Printf("%v\n", c)
  c.Put(4, "str4") // {1: “str1”, 3: “str2”, 4: “str4”}
  fmt.Println("c.Put(4) = str4\n")
  fmt.Printf("%v\n", c)
}
