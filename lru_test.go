package main

import (
  "fmt"
  "math/rand"
  "testing"
)

func TestEviction(t *testing.T) {
  c := NewCache(3) // {}
  c.Put(1, "str1") // {1: “str1”}
  c.Put(2, "str2") // {1: “str1”, 2: “str2”}
  c.Put(3, "str3") // {1: “str1”, 2: “str2”, 3: “str3”}
  c.Get(3)
  c.Get(2)
  c.Get(1)
  c.Get(3)
  c.Put(4, "str4") // {1: “str1”, 3: “str2”, 4: “str4”}
  rightOrder := []uint32{1, 3, 4}
  i := 0
  for node := c.ll.Front(); node != nil; node = node.Next() {
    if node.Value.(*Data).Key != rightOrder[i] {
      t.Errorf("wrong order")
      t.Fail()
      return
    }
    i += 1
  }
}

func TestZeroSize(t *testing.T) {
  const size = 0
  c := NewCache(size)
  if c != nil {
    c.Put(2, "hello 2")
  }
}

func TestOrder(t *testing.T) {
  const size = 100
  c := NewCache(size)
  ar := []uint32{}
  var i uint32
  for i = 0; i < size; i++ {
    c.Put(i, fmt.Sprintf("key %d", i))
    ar = append(ar, i)
  }

  rand.Shuffle(len(ar), func(i, j int) {
    ar[i], ar[j] = ar[j], ar[i]
  })

  for _, i = range ar {
    if rand.Float32() > 0.5 {
      c.Get(i)
    } else {
      c.Put(i, fmt.Sprintf("new key %d", i))
    }
  }

  node := c.ll.Front()
  for i = 0; i < size; i++ {
    if node.Value.(*Data).Key != ar[i] {
      t.Errorf("wrong order")
      t.Fail()
      return
    }
    node = node.Next()
  }
}