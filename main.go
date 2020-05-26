package main

import (
	"fmt"
	"go-algo/cache"
)

func main() {
	lru := cache.New(3)

	lru.Add("A", 1)
	lru.Add("B", 2)
	lru.Add("C", 3)
	lru.Add("D", 4)
	lru.Add("E", 5)

	fmt.Printf("lru %v ", lru)
}
