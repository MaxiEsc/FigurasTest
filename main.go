package main

import (
	"fmt"
	"paquetes/figuras"
)

func main() {

	//Luego de completarlos... esto es innecesario en fin es otra forma de hacerlo
	var ancho float32
	var alto float32
	var radio float32
	var pi float32

	//Ingresamos los numero por teclado
	fmt.Println("Ingrese el ancho del cuadrado")
	fmt.Scanln(&ancho)

	fmt.Println("Ingrese el alto del cuadrado")
	fmt.Scanln(&alto)

	fmt.Println("Ingrese el radio de circulo")
	fmt.Scanln(&radio)

	fmt.Println("Ingrese el pi de circulo")
	fmt.Scanln(&pi)

	//Ahora utilizamos la interfaz correspondiente
	//Creamos figuras de prueba sin datos pues ya se los cargamos...
	//La verdad es que podriamos haberlos creado con la figura sin lanecesidad de
	//pasarlos por separados y sacarlos desde la misma figura
	cua := figuras.Cuadrado{}
	cir := figuras.Circulo{}

	//Siempre pasar las figuras mediante referencia... NO olvidar!!!! si no como las detecta
	//Si hay 2 perimetros... y dos areas... XD
	figuras.CalculaArea(&cua, ancho, alto)
	figuras.CalcularPerimetro(&cua, ancho, alto)

	figuras.CalculaArea(&cir, pi, radio)
	figuras.CalcularPerimetro(&cir, pi, radio)
}
