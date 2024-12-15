package main

import "fmt"

const pontoEbulicaoKelvin = 373.0

func main() {
	temperaturaCelsius := pontoEbulicaoKelvin - 273.0

	fmt.Printf("O ponto de ebulição da água %g° em Kelvin é representada como %g° em Celsius", pontoEbulicaoKelvin, temperaturaCelsius)
}
