package main

import "fmt"

func main() {

	license := false
	age := 19

	// tarea a realizar
	if license == true && age >= 15 {
		fmt.Println("Puedes seguir conduciendo")
	} else if license == false && age >= 15 {
		fmt.Println("No tienes licencia, no puedes conducir")
	} else {
		fmt.Println("No tienes la edad y licencia para conducir")
	}

}
