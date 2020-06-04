package main

import (
	"fmt"
	"go-algo/fb"
)

func main() {

	cache := fb.NewCache(3)

	fmt.Printf("LRUGet %v \n", cache.LRUGet(1))

	fmt.Printf("LRUPut %v \n", cache.LRUPut(1, "A"))
	fmt.Printf("LRUGet %v \n", cache.LRUGet(1))

	fmt.Printf("LRUPut %v \n", cache.LRUPut(2, "B"))
	fmt.Printf("LRUPut %v \n", cache.LRUPut(3, "C"))
	fmt.Printf("LRUPut %v \n", cache.LRUPut(4, "D"))

	fmt.Printf("LRUGet %v \n", cache.LRUGet(1))
	fmt.Printf("LRUGet %v \n", cache.LRUGet(2))
	//fmt.Printf("LRUGet %v \n", cache.LRUGet(1))

}
