package main

import (
  "container/list"
  "fmt"
)

const DataSize = 128

type table map[uint32]*list.Element

type Cache struct {
  ll     *list.List
  ht     table
  size   int
  memory int
}

type Data struct {
  Key   uint32
  Value string
}

func NewCache(size int) *Cache {
  if size <= 0 {
    return nil
  }
  return &Cache{list.New(), make(table), size, 0}
}

func (c *Cache) SetMemory(memory int) {
  c.memory = memory
}

func (c *Cache) String() string {
  ar := []string{}
  for i := c.ll.Front(); i != nil; i = i.Next() {
    ar = append(ar, fmt.Sprintf("%p:%v", i, *(i.Value.(*Data))))
  }
  return fmt.Sprintf("List: %v\nHT: %v\n", ar, c.ht)
}

func (c *Cache) Put(key uint32, val string) {
  if el, found := c.ht[key]; !found {
    // no such key exists
    // insert new key
    if c.ll.Len() >= c.size || (c.memory > 0 && c.ll.Len()*DataSize >= c.memory) {
      // LRU eviction
      data := c.ll.Remove(c.ll.Front()).(*Data)
      delete(c.ht, data.Key)
    }
    newEl := c.ll.PushBack(&Data{key, val})
    c.ht[key] = newEl
  } else {
    // cache entry exists, update it
    c.ll.MoveToBack(el)
    el.Value.(*Data).Value = val
  }
}

func (c *Cache) Get(key uint32) string {
  if el, found := c.ht[key]; !found {
    // no such key exists
    return ""
  } else {
    // key exists
    c.ll.MoveToBack(el)
    return el.Value.(*Data).Value
  }
}
