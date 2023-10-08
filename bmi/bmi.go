package bmi

import (
	"fmt"
)

/*
How much do you weigh? (don't lie)
80
How tall are you? (barefoot)
1.85
Right now your BMI is 23.37
You have a normal weight, I have healthy envy of you

Use the following ranges for the final message:
- If BMI is less than 18.5 -> 'You are underweight, add more potato to the broth'
- If BMI is greater than or equal to 18.5 but less than 25 -> 'You have a normal weight, I have healthy envy of you'
- If BMI is greater than or equal to 25 -> 'You are overweight, I know, the pandemic has affected us all'
Assume that the user will input valid numbers.
*/

func GetBmi() {
	fmt.Println("¿Cuanto pesas en kilogramos? (no mientas). Ejemplo 80: ")
	var weight float64
	fmt.Scan(&weight)
	fmt.Println("¿Cual es tu altura en metros? (descalzo). Ejemplo 1.85 metros: ")
	var height float64
	fmt.Scan(&height)

	bmi := calculateBmi(weight, height)
	fmt.Println("Ahora mismo tu IMC es de: ", bmi)

	if bmi < 18.5 {
		fmt.Println("Estás bajo de peso, añade más papa al caldo")
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Println("Tienes un peso normal, te tengo envidia sana")
	} else if bmi >= 25 {
		fmt.Println("Tienes sobrepeso, lo sé, la pandemia nos ha afectado a todos")
	}
}

func calculateBmi(weight float64, height float64) float64 {
	var bmi float64 = weight / (height * height)
	return bmi
}
