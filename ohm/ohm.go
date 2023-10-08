package ohm

import (
	"fmt"
)

/*
v = i * r
i = v / r
r = v / i

func ohm(v, r, i){...}

print(ohm(0, 2, 3)) // 6 -> existen r, i -> OK
print(ohm(6, 2)) // 3 -> existen v, r -> OK
print(ohm(6, 0, 3)) // 2 -> existen v, i -> OK
print(ohm(1, 2, 3)) // 0  -> existen v, r, i -> OK
print(ohm(0, 0, 3)) // 0  -> cualquier otro caso -> OK
*/

func OhmLaw() {
	//fmt.Println("Ingresa los valores para 'voltaje', 'resistencia' e 'intensidad' en ese mismo orden: ")
	//fmt.Scan(&v, &r, &i)
	var v float32
	var r float32
	var i float32
	var resultMessage string = "El resultado es: "

	fmt.Println("Ingresa el valor para 'voltaje': ")
	fmt.Scan(&v)
	fmt.Println("Ingresa el valor para 'resistencia': ")
	fmt.Scan(&r)
	fmt.Println("Ingresa el valor para 'intensidad': ")
	fmt.Scan(&i)

	if v > 0 && r > 0 && i > 0 { //-> existen v, r, i
		print(resultMessage, 0)
	} else if v > 0 && r > 0 { //-> existen v, r
		print(resultMessage, "Intensidad i =", v/r)
	} else if v > 0 && i > 0 { //-> existen v, i
		print(resultMessage, "Resistencia r =", v/i)
	} else if r > 0 && i > 0 { //-> existen r, i
		print(resultMessage, "Voltaje v =", i*r)
	} else {
		print(resultMessage, 0) //-> cualquier otro caso
	}
}
