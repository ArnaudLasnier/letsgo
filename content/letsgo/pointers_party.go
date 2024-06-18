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

func (backMaker BackMaker) PartyWithoutPointer() {
	backMaker.Age++
}

func main() {
	// START OMIT
	arnaud := BackMaker{Name: "Arnaud", Age: 29}
	arnaud.PartyWithoutPointer()
	fmt.Println("without pointer:", arnaud.Age)
	arnaud.Party()
	fmt.Println("with pointer:", arnaud.Age)
	// END OMIT
}
