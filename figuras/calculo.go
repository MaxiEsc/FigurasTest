package figuras

import "fmt"

//Uso de interface
type Calculo interface {
	perimetro(a float32, b float32) string
	area(a float32, b float32) string
}

//Metodo de la interfaz
func CalcularPerimetro(calculo Calculo, a float32, b float32) {
	fmt.Println(calculo.perimetro(a, b))
}

func CalculaArea(calculo Calculo, a float32, b float32) {
	fmt.Println(calculo.area(a, b))
}
