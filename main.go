package main

import (
	"fmt"
	"go-algo/general"
)

func main() {

	tb := general.NewLeakyBucket(3, 10)

	fmt.Printf("result %v %v\n", 1, tb.IsAllowed("A", 1))
	fmt.Printf("result %v %v\n", 2, tb.IsAllowed("A", 2))
	fmt.Printf("result %v %v\n", 2, tb.IsAllowed("A", 2))
	fmt.Printf("result %v %v\n", 4, tb.IsAllowed("A", 4))
	fmt.Printf("result %v %v\n", 10, tb.IsAllowed("A", 10))
	fmt.Printf("result %v %v\n", 12, tb.IsAllowed("A", 12))

	fmt.Printf("result %v %v\n", "", tb)
	tb.Cleanup(20)
	fmt.Printf("result %v %v\n", "", tb)
}
