package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		base   float64
		altura float64
	)
	fmt.Print("Vamos a calcular la hipotenusa, area y el perímetro inserte el valor de la base del triangulo:")
	fmt.Scan(&base)
	fmt.Print("Inserte el valor de la altura del triangulo:")
	fmt.Scan(&altura)

	hipotenusa := math.Sqrt(math.Pow(base, 2) + math.Pow(altura, 2))
	perimetro := base + altura + hipotenusa
	area := ((base) * (altura)) / 2

	fmt.Printf("La hipotenusa es: %.3v \n", hipotenusa)
	fmt.Printf("El Area es: %v \n ", area)
	fmt.Printf("El perímetro es: %.4v \n ", perimetro)

}
