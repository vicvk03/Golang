// Ejercicio para aprender sobre variables en el lenguaje Go.
// En este lenguaje siempre se tienen que declarar las variables antes de usarlas.

package main

import (
	"fmt"
	"strconv" // Importamos strconv para realizar conversiones de String a otros tipos de datos
	"unsafe"  // Importamos unsafe para revisar la dirección de memoria de las variables
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

	// Tipos de enteros (Int) dentro de Go.
	// Generalmente, se asigna la longitud del dato, de acuerdo al sistema operativo. Como se describe en la siguiente tabla:
	// Lo mismo aplica para los tipos de enteros sin signo (uint), a excepción que únicamente acepta valor reales (sin signo).
	/*
	 * int8   - 8 bits (1 byte) - Rango: -128 a 127
	 * int16  - 16 bits (2 bytes) - Rango: -32,768 a 32,767
	 * int32  - 32 bits (4 bytes) - Rango: -2,147,483,648 a 2,147,483,647
	 * int64  - 64 bits (8 bytes) - Rango: -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807
	 */

	// Impresión de línea con formato

	fmt.Println() // Salto de línea

	var myInt8Var int8
	fmt.Printf("El valor de la variable entera 8 bits, es: %d\n", myInt8Var)

	// Impresión de la cantidad de memoria instanciada la variable entera 8 bits
	fmt.Printf("La cantidad de memoria ocupada por la variable entera 8 bits, es: %d bytes\n", unsafe.Sizeof(myInt8Var))        // Dada por Copilot
	fmt.Printf("Tipo_Dato: %T, Alm_Bytes: %d, Alm_Bits: %d\n", myInt8Var, unsafe.Sizeof(myInt8Var), unsafe.Sizeof(myInt8Var)*8) // Dada por el curso

	var myInt16Var int16
	fmt.Printf("Tipo_Dato: %T, Alm_Bytes: %d, Alm_Bits: %d\n", myInt16Var, unsafe.Sizeof(myInt16Var), unsafe.Sizeof(myInt16Var)*8) // Dada por el curso

	var myInt32Var int32
	fmt.Printf("Tipo_Dato: %T, Alm_Bytes: %d, Alm_Bits: %d\n", myInt32Var, unsafe.Sizeof(myInt32Var), unsafe.Sizeof(myInt32Var)*8) // Dada por el curso

	var myInt64Var int64
	fmt.Printf("Tipo_Dato: %T, Alm_Bytes: %d, Alm_Bits: %d\n", myInt64Var, unsafe.Sizeof(myInt64Var), unsafe.Sizeof(myInt64Var)*8) // Dada por el curso

	var varInt int
	fmt.Printf("Tipo_Dato: %T, Alm_Bytes: %d, Alm_Bits: %d\n", varInt, unsafe.Sizeof(varInt), unsafe.Sizeof(varInt)*8) // Dada por el curso

	// Tipos de datos flotantes (float) dentro de Go.
	var myFloat32Var float32
	fmt.Printf("El valor de la variable flotante 32 bits, es: %f\n", myFloat32Var)
	fmt.Printf("Tipo_Dato: %T, Alm_Bytes: %d, Alm_Bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8) // Dada por el curso

	var myFloat64Var float64
	fmt.Printf("El valor de la variable flotante 64 bits, es: %f\n", myFloat64Var)
	fmt.Printf("Tipo_Dato: %T, Alm_Bytes: %d, Alm_Bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8) // Dada por el curso

	// Strings en Go
	var myStringVar3 string
	myStringVar3 = "Hola, soy un string 3"
	fmt.Printf("El valor de la variable string 3, es: %s\n", myStringVar3)

	// Strings con más de una línea
	var myMultilineString string = `Hola, soy un string
			con varias líneas.`
	fmt.Println("El valor de la variable string multilinea, es:", myMultilineString)

	// Conversiones (aisladas) del código principal
	{
		// Declaración de una variable de tipo flotante
		fmt.Println()
		floatVar := 33.11
		fmt.Printf("Tipo_Dato: %T, Valor_Variable: %f\n", floatVar, floatVar)

		// Conversión de float a String
		floatStrVar := fmt.Sprintf("%.2f", floatVar) // Usamos Sprintf para formatear el número a String, mostrando únicamente dos decimales originales de la variable float.
		fmt.Printf("Tipo_Dato: %T, Valor_Variable: %s\n", floatStrVar, floatStrVar)

		// Conversión de Int a String
		intVar := 42
		fmt.Printf("Tipo_Dato: %T, Valor_Variable: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("Tipo_Dato: %T, Valor_Variable: %s\n", intStrVar, intStrVar)

		intVal1, err := strconv.ParseInt("13359", 10, 64) // Convertimos el String a Int, con base 10 y tamaño de 64 bits
		fmt.Println("Error al convertir de String a Int:", err)
		fmt.Printf("Tipo_Dato: %T, Valor_Variable: %d\n", intVal1, intVal1)

		intVal2, err := strconv.ParseInt("fa123", 10, 64)
		fmt.Println("Error al convertir de String a Int:", err)
		fmt.Printf("Tipo_Dato: %T, Valor_Variable: %d\n", intVal2, intVal2)

		floatVar1, _ := strconv.ParseFloat("68.29", 64) // Convertimos el String a Float, con tamaño de 64 bits / Con el guion bajo "_", se descarta el error.
		fmt.Printf("Tipo_Dato: %T, Valor_Variable: %f\n", floatVar1, floatVar1)

	}

}
