package main

import (
  "container/list"
  "fmt"
  "time"
)

const DataSize = 128

type table map[uint32]*list.Element

type Cache struct {
  ll     *list.List
  ht     table
  size   int
  memory int
  ttl    time.Duration
}

type Data struct {
  Key       uint32
  Value     string
  Timestamp time.Time
}

func NewCache(size int) *Cache {
  if size <= 0 {
    return nil
  }
  return &Cache{list.New(), make(table), size, 0, 0}
}

func (c *Cache) SetMemory(memory int) {
  c.memory = memory
}

func (c *Cache) SetTTL(ttl int) {
  c.ttl = time.Duration(ttl) * time.Second
}

func (c *Cache) String() string {
  ar := []string{}
  for i := c.ll.Front(); i != nil; i = i.Next() {
    data := i.Value.(*Data)
    ar = append(ar, fmt.Sprintf("%p:(%d, %s)", i, data.Key, data.Value))
  }
  return fmt.Sprintf("List: %v\nHT: %v\n", ar, c.ht)
}

func (c *Cache) delete(el *list.Element) {
  data := c.ll.Remove(el).(*Data)
  delete(c.ht, data.Key)
}

func (c *Cache) Put(key uint32, val string) {

  if el, found := c.ht[key]; !found {
    // no such key exists
    // insert new key
    if c.ll.Len() >= c.size || (c.memory > 0 && c.ll.Len()*DataSize >= c.memory) {
      // LRU eviction
      c.delete(c.ll.Front())
    }
    data := &Data{
      Key:   key,
      Value: val,
    }
    newEl := c.ll.PushBack(data)
    if c.ttl > 0 {
      data.Timestamp = time.Now()
    }
    c.ht[key] = newEl
  } else {
    // cache entry exists, update it
    c.ll.MoveToBack(el)
    el.Value.(*Data).Value = val
    if c.ttl > 0 {
      el.Value.(*Data).Timestamp = time.Now()
    }
  }
}

func (c *Cache) Get(key uint32) string {
  if el, found := c.ht[key]; !found {
    // no such key exists
    return ""
  } else {
    // key exists
    // check ttl
    if c.ttl > 0 && time.Since(el.Value.(*Data).Timestamp) > c.ttl {
      c.delete(el)
      return "not found"
    }
    c.ll.MoveToBack(el)
    return el.Value.(*Data).Value
  }
}
