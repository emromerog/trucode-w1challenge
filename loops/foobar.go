package loops

import (
	"fmt"
	"strconv"
)

/*
Para los múltiplos de tres imprima "Foo" en lugar del número.
Para los múltiplos de cinco imprima "Bar".
Para los que son múltiplos de tres y cinco, escriba "FooBar".

print(foobar(3)) // 1->2->Foo
print(foobar(5)) // 1->2->Foo->4->Bar
print(foobar(15)) //1->2->Foo->4->Bar->Foo->7->8->Foo->Bar->11->Foo->13->14->FooBar
*/

func Foobar() {
	fmt.Println("Ingresa el número: ")
	var num int
	fmt.Scan(&num)
	var chainstr string

	for i := 1; i <= num; i++ {
		if i == num {
			chainstr += parseNum(i)
		} else {
			chainstr += parseNum(i) + "->"
		}
	}

	print(chainstr)
}

func parseNum(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "FooBar"
	} else if num%3 == 0 {
		return "Foo"
	} else if num%5 == 0 {
		return "Bar"
	} else {
		return strconv.Itoa(num) // convertir entero num a string para retornarlo.
	}
}
