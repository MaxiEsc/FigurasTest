package figuras

import "fmt"

type Cuadrado struct {
	ancho  float32
	altura float32
}

func (*Cuadrado) area(a float32, b float32) string {
	return fmt.Sprintf("El Área del Cuadrado es: %v", a*b)
}

func (*Cuadrado) perimetro(a float32, b float32) string {
	return fmt.Sprintf("El Perimetro del Cuadrado es: %v", (2*a + 2*b))
}