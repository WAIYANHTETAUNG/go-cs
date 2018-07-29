package main

import (
	"fmt"

	"github.com/pyaesone17/go-cs/hashmap"
)

func main() {
	hashtable := hashmap.NewHashTable()
	fmt.Println(hashtable.Get("veve"))
}
