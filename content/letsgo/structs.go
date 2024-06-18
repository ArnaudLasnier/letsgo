package main

import (
	"fmt"
)

type BackMaker struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (backMaker *BackMaker) Party() {
	backMaker.Age++
}

func main() {
	arnaud := BackMaker{Name: "Arnaud", Age: 29}
	arnaud.Party()
	fmt.Println(arnaud.Age)
}
