package main

import (
	"fmt"
)

type HashTable struct {
	MAX int
	arr []interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{
		MAX: 100,
		arr: make([]interface{}, 100),
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
	ht.arr[h] = value
}

func (ht *HashTable) GetItem(key string) interface{} {
	h := ht.getHash(key)
	return ht.arr[h]
}

func (ht *HashTable) DeleteItem(key string) {
	h := ht.getHash(key)
	ht.arr[h] = nil
}

func main() {
	ht := NewHashTable()
	ht.SetItem("key1", 42)
	ht.SetItem("key2", "value2")

	fmt.Println("Value for key1:", ht.GetItem("key1")) // Output: 42
	fmt.Println("Value for key2:", ht.GetItem("key2")) // Output: value2

	ht.DeleteItem("key1")
	fmt.Println("Value for key1 after deletion:", ht.GetItem("key1")) // Output: nil
}
