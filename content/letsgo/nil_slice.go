package main

import "fmt"

func main() {
	// START OMIT
	var nilSlice []int

	fmt.Println(len(nilSlice))
	fmt.Println(cap(nilSlice))
	fmt.Println(nilSlice == nil)
	// END OMIT
}
