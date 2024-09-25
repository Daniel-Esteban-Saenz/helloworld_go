package main

import "fmt"

func dvision_residuo() {
	var x int
	var y int
	fmt.Println("inserte el valor de X")
	fmt.Scanln(&x)
	fmt.Println("inserte el valor de Y")
	fmt.Scanln(&y)
	fmt.Println("la division de x y y es:", x/y)
	fmt.Println("el residuo de x/y es: ", x % y)
}


func par_impar() {
	var x int
	fmt.Println("inserte el valor de X")
	fmt.Scanln(&x)
	if x%2 == 0 {
        fmt.Println(x, "es par")
    } else {
        fmt.Println(x, "es impar")
    }
}

func main() {
	// dvision_residuo()
	par_impar()
}