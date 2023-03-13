package figuras

import "fmt"

type Circulo struct {
	radio float32
	pi    float32
}

func (*Circulo) area(a float32, b float32) string {
	return fmt.Sprintf("El Área del circulo es: %v", (b * (a * a)))
}

func (*Circulo) perimetro(a float32, b float32) string {
	return fmt.Sprintf("El Perimetro del Cuadrado es: %v", (2 * b * a))
}