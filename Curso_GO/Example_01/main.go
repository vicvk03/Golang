// Ejercicio para aprender sobre variables en el lenguaje Go.
// En este lenguaje siempre se tienen que declarar las variables antes de usarlas.

package main

import (
	"fmt"
)

func main() {

	// Declaración de variables de diferentes tipos de datos
	// Tipo de dato entero (int)
	var myIntVar int
	myIntVar = -12

	fmt.Println("El valor de la variable entera, es:", myIntVar)

	// Tipo de dato entero sin signo (uint)
	var myUintVar uint
	myUintVar = 12
	fmt.Println("El valor de la variable u-entera, es:", myUintVar)

	// Tipo de dato String
	var myStringVar string
	myStringVar = "Hola, soy un string"
	fmt.Println("El valor de la variable string, es:", myStringVar)

	// Tipo de dato Booleano
	var myBoolVar bool
	myBoolVar = true
	fmt.Println("El valor de la variable booleana, es:", myBoolVar)

	// Revisión de dirección de memoria de variables
	fmt.Println("Dirección de memoria de string:", &myStringVar)

	// Asignación de valores a variables sin declararlas previamente (implícitas)
	myIntVar2 := 100
	fmt.Println("El valor de la variable entera 2, es:", myIntVar2)

	myStringVar2 := "Hola, soy otro string"
	fmt.Println("El valor de la variable string 2, es:", myStringVar2)

	// Declaración de constantes
	// Las constantes pueden tomar el valor (indiferente e implícitamente) que se les asigne y no podrán ser modificadas
	const myFirstConst = 42
	const mySecondConst = "Soy una constante"

	fmt.Println("El valor de la constante 1, es:", myFirstConst)
	fmt.Println("El valor de la constante 2, es:", mySecondConst)

	const myIntConst = 1000
	fmt.Println("El valor de la constante entera, es:", myIntConst)

	// (Intento de) Modificar una constante
	// mySecondConst = "Nuevo valor"                                            // Esto generará un error de compilación, ya que las constantes no pueden ser modificadas
	// fmt.Println("El valor de la constante 2 modificado, es:", mySecondConst) // Esto no se ejecutará debido a lo expuesto anteriormente

}
