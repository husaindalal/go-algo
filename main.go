package main

import (
	"fmt"
	"go-algo/fb"
)

func main() {

	result := fb.ArrangeLetters([]string{
		"eterq",
		"errsd",
		"yrtwe",
		"yrzdvs",
		"fhhs",
		"fff",
		"tjsg",
		"tjadsd",
		"tjadsdrr",
	})

	for _, r := range result {
		fmt.Printf("arranged %c\n", r)

	}
}
