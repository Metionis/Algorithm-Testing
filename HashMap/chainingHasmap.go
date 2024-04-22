package main

import (
    "fmt"
)

type HashTable struct {
    MAX  int
    arr  [][][2]interface{}
}

func NewHashTable() *HashTable {
    return &HashTable{
        MAX: 100,
        arr: make([][][2]interface{}, 100),
    }
}

func (ht *HashTable) getHash(key string) int {
    h := 0
    for _, char := range key {
        h += int(char)
    }
    return h % ht.MAX
}

func (ht *HashTable) SetItem(key string, value interface{}) {
    h := ht.getHash(key)
    found := false
    for idx, element := range ht.arr[h] {
        if len(element) == 2 && element[0].(string) == key {
            ht.arr[h][idx] = [2]interface{}{key, value}
            found = true
            break
        }
    }
    if !found {
        ht.arr[h] = append(ht.arr[h], [2]interface{}{key, value})
    }
}

func (ht *HashTable) GetItem(key string) interface{} {
    h := ht.getHash(key)
    for _, element := range ht.arr[h] {
        if element[0].(string) == key {
            return element[1]
        }
    }
    return nil
}

func (ht *HashTable) DeleteItem(key string) {
    h := ht.getHash(key)
    for index, element := range ht.arr[h] {
        if element[0].(string) == key {
            ht.arr[h] = append(ht.arr[h][:index], ht.arr[h][index+1:]...)
            break
        }
    }
}

func main() {
    ht := NewHashTable()
    ht.SetItem("name", "John")
    ht.SetItem("age", 30)
    fmt.Println(ht.GetItem("name")) // Output: John
    ht.DeleteItem("name")
    fmt.Println(ht.GetItem("name")) // Output: <nil>
}
