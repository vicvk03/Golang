package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("Tipo: %T, bytes: %d, bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	// Vectores
	var myArrayVar [6]int
	fmt.Println(myArrayVar)

	myArrayVar1 := [3]string{"valor1", "valor2", "valor3"}
	fmt.Println(myArrayVar1)

}
