package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func imprimir() {
	fmt.Println("texto desde la funcion imprimir")
}

func hola_string(s string) string {
	return "Hola " + s
}

func suma_int(a int, b int) int {
	return a + b
}

func comparar_bool() {
	var a bool
	b := true
	a = false
	if a == b {
		fmt.Println("a and b are same")
	} else {
		fmt.Println("a and b arent same")
	}

}

func arreglo() {
	var aprendices [5]string
	aprendices[0] = "Leonardo"
	aprendices[1] = "Juan Sebastian"
	aprendices[2] = "Franck"
	aprendices[3] = "Juan jose"
	aprendices[4] = "Jaider"

	fmt.Println(aprendices[1])
}

func tipos_datos() {
	var Texto string = "Sebastian"
	var numero int = 1000
	var decimal float64 = 0.00001

	fmt.Println(reflect.TypeOf(Texto))
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.TypeOf(decimal))
}

func conver_strig_to_boolean() {
	var palabra string = "true"
	boolean, err := strconv.ParseBool(palabra)
	fmt.Println(boolean, reflect.TypeOf(boolean))
	fmt.Println("este es el error", err)
}

func conver_bool_to_string() {
	var palabra_bool bool = true
	strbool := strconv.FormatBool(palabra_bool)
	fmt.Println(strbool, reflect.TypeOf(palabra_bool))
}

func variables() {
	var nombre, apellido string = "Daniel", "Sáenz"
	fmt.Println(nombre, apellido)
}

func variables_bloque() {
	var (
		nombre     string = "Daniel"
		apellido   string = "Sáenz"
		edad       int    = 20
		estudiando bool   = true
	)
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Apellido: ", apellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Estudiando: ", estudiando)
}

func valor_cero() {
	var nombre_cero string
	var edad_cero int
	var peso_cero float64
	var estudiante_cero bool
	fmt.Println("Nombre: ", nombre_cero)
	fmt.Println("Edad: ", edad_cero)
	fmt.Println("Peso: ", peso_cero)
	fmt.Println("Estudiante: ", estudiante_cero)
}

func declaracion_corta() {
	Nombre := "Daniel Sáenz"
	edad := 20
	peso := 59
	estudiante := true
	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Edad: ", edad)
	fmt.Println("Peso: ", peso)
	fmt.Println("Estudiante: ", estudiante)
}

func declaracion_global() {
	var profesion = "Deportista"
	sueldo := 1000000
	fmt.Println("Profesion: ", profesion)
	fmt.Println("Sueldo: ", sueldo)
}

func scope() {
	var var1 = "este es el nivel 1"
	var var2 = "este es el nivel 2"
	{
		var var3= "este es el nivel 3"
		fmt.Println(var3)
	}
	fmt.Println(var1)
	fmt.Println(var2)
}

func puntero() {
	vehiculo1 := "rojo"
    fmt.Println("El vehiculo1 es", vehiculo1)

    vehiculo2 := vehiculo1
    fmt.Println("El vehiculo2 es", vehiculo2)

    vehiculo3 := &vehiculo1
    fmt.Println("El vehiculo3 es", *vehiculo3)	

    vehiculo1 = "gris"

    fmt.Println("El vehiculo1 es", vehiculo1, &vehiculo1)
    fmt.Println("El vehiculo2 es", vehiculo2, &vehiculo2)
    fmt.Println("El vehiculo3 es", *vehiculo3, vehiculo3)
}

func conversionEnPies(altura *float32) float32 {
    *altura = *altura - 0.10
    return *altura
}

var altura float32 = 1.70

const Pi = 3.1416

func area(radio float64) float64 {
    return Pi * radio * radio
}


func main() {

	// fmt.Println("Texto desde el main")
	// imprimir()
	// fmt.Println(hola_string("jaime"))
	// fmt.Println(suma_int(3, 5))
	// comparar_bool()
	// arreglo()
	// tipos_datos()
	// conver_strig_to_boolean()
	// conver_bool_to_string()
	// variables()
	// variables_bloque()
	// valor_cero()
	// declaracion_corta()
	// declaracion_global()
	// scope()
	puntero()
	// fmt.Println("La altura es:", altura, "mts")
	// fmt.Println("Al envejecer:", conversionEnPies(&altura), "mts")
	// fmt.Println("Despues de envejecer:", altura, "mts")
	fmt.Println("El area de un circulo cuyo radio es 3 es: ", area(3))
}
