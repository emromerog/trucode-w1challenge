package evenodd

import (
	"fmt"
)

func Evenodd() {
	fmt.Println("Ingresa el número que deseas saber si es par o impar: ")
	var num int
	fmt.Scan(&num)
	if num%2 == 0 {
		print("even")
	} else {
		print("odd")
	}
}
