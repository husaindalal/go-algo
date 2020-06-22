package main

import (
	"fmt"
	"go-algo/hash"
)

func main() {


	result := hash.ReplaceWords([]string{"car", "cat", "rat", "bat"}, "the cattle was rattled by the battery")

	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

}

