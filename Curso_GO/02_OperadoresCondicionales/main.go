package main

import (
	"fmt"
)

func main() {

	// Operadores condicionales
	// ==, !=, >, <, >=, <=
	fmt.Println("Operadores Condicionales")

	// Comparan valores y devuelven un booleano
	a := 10
	b := 20

	fmt.Println("a == b:", a == b) // false
	fmt.Println("a != b:", a != b) // true
	fmt.Println("a > b:", a > b)   // false
	fmt.Println("a < b:", a < b)   // true
	fmt.Println("a >= b:", a >= b) // false
	fmt.Println("a <= b:", a <= b) // true

	// salto de linea
	fmt.Println()

	// Operadores lógicos
	// '&&' (AND), '||' (OR), '!' (DISTINCT)
	fmt.Println("Operadores Lógicos")

	// Evalúa valores distintos y devuelve un booleano
	c := true
	d := false

	fmt.Println("c && d:", c && d) // false
	fmt.Println("c || d:", c || d) // true
	fmt.Println("!c:", !c)         // false
	fmt.Println("!d:", !d)         // true

	// salto de linea
	fmt.Println()

	// Condicionales

	// if
	yearsOld := 20

	if yearsOld > 18 {
		fmt.Printf("Tienes %d, eres mayor de edad!\n", yearsOld)
	}

	// salto de linea
	fmt.Println()

	boolVal := true
	if boolVal {
		fmt.Println("El valor es verdadero")
	} else {
		fmt.Println("El valor es falso")
	}

	if value := true; value == true {
		fmt.Println("El valor es verdadero")
	}

	// salto de linea
	fmt.Println()

	// switch

	numbreVar := 2

	switch numbreVar {
	case 1:
		fmt.Println("El número es uno")
	case 2:
		fmt.Println("El número es dos")
	case 3:
		fmt.Println("El número es tres")
	default:
		fmt.Println("El número no es ni uno, ni dos, ni tres")
	}

}
