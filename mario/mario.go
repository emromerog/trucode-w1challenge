package mario

import (
	"fmt"
)

func Mario() {
	var height int
	appAlive := true

	//Solicitar la altura de la piramide hasta que el usuario ingrese un numero entre 1 y 8.
	for appAlive {
		fmt.Print("Ingrese la altura de la piramide (un valor entre 1 y 8): ")
		fmt.Scan(&height)
		if height < 1 || height > 8 {
			fmt.Println("La altura debe estar entre 1 y 8.")
		} else {
			appAlive = false
		}
		fmt.Println("Altura de la piramide: ", height)
	}

	for i := 0; i < height; i++ {
		// Imprimir espacios en blanco hasta que la cantidad de espacios sea igual a la altura menos 1.
		for j := 0; j < height-i-1; j++ {
			fmt.Print(" ")
		}
		// Imprimir bloques #
		for k := 0; k <= i; k++ {
			fmt.Print("#")
		}

		// Ir a la siguiente línea después de imprimir cada nivel de la piramide, es decir, hacer un salto de linea.
		fmt.Println()
	}
}
