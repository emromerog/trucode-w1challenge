# Functions challenge

## Even or Odd

Crea una función que tome un número como argumento y verifica en ella si ese número es par o impar.

```go
print(evenOrOdd(1)) // "odd"
print(evenOrOdd(22)) // "even"
```

## Ley de Ohm

Cree una función que calcule el valor faltante de 3 entradas utilizando la ley de Ohm. Los argumentos son v, r o i (también conocidas como voltaje, resistencia y corriente).

La ley de Ohm dice que: "la intensidad de la corriente eléctrica que circula por un conductor eléctrico es directamente proporcional al voltaje aplicado e inversamente proporcional a la resistencia del mismo".

La sentencia anterior nos lleva a determinar que:

- $V = IR$
- $I = V/R$
- $R = V/I$

Donde $V$ es voltaje, $I$ es intensidad y $R$ es resistencia.

```go
func ohm(v, r, i){...}

print(ohm(0, 2, 3)) // 6
print(ohm(6, 2)) // 3
print(ohm(6, 0, 3)) // 2
print(ohm(1, 2, 3)) // 0
print(ohm(0, 0, 3)) // 0
```

## Foobar

Escriba una función que imprima los números indicados en el argumento num.

```go
func foobar(num){...}
```

Para los múltiplos de tres imprima "Foo" en lugar del número y para los múltiplos de cinco imprima "Bar". Para los que son múltiplos de tres y cinco, escriba "FooBar".

```go
print(foobar(3)) // 1->2->Foo
print(foobar(5)) // 1->2->Foo->4->Bar
print(foobar(15)) //1->2->Foo->4->Bar->Foo->7->8->Foo->Bar->11->Foo->13->14->FooBar
```