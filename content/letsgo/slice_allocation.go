package main

import "fmt"

func main() {
	// START OMIT
	slice := make([]int, 10)
	slice[1] = 1
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	// END OMIT
}
