package main

import (
	"os"
	"w1challenge/bmi"
	"w1challenge/evenodd"
	"w1challenge/loops"
	"w1challenge/mario"
	physics "w1challenge/ohm"
)

func main() {
	//Solicitar el argumento que se desea ejecutar.
	args := os.Args
	option := args[1]

	switch option {
	case "evenodd":
		evenodd.Evenodd()
	case "ohm":
		physics.OhmLaw()
	case "foobar":
		loops.Foobar()
	case "bmi":
		bmi.GetBmi()
	case "mario":
		mario.Mario()
	default:
		print("Las opciones válidas son: evenodd, ohm, foobar, bmi, mario")
	}
}
