package main

import "fmt"

func main() {

	//Declaracion de una Costante
	const num float64 = 3.14
	const num2 = 3.1415

	fmt.Println("Numero 1:", num)
	fmt.Println("Numero 2:", num2)

	//Declaracion de variables enteras --> Si no estan siendo usadas saltara un error
	num3 := 5
	var num4 = 6
	var num5 int

	fmt.Println(num3, num4, num5)

	//Zero values
	var a int
	var b string
	var c float64
	var d bool

	fmt.Println(a, b, c, d)

	//Variables Constantes
	const baseCuadrada = 10
	areaCuadrada := baseCuadrada * baseCuadrada
	fmt.Println("Area Cuadrada: ", areaCuadrada)
}
