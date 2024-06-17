package main

import "fmt"

func main() {
	var slice []int
	for i := range 100 {
		slice = append(slice, i)
	}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
