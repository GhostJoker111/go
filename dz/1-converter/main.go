package main

import "fmt"

const usdToEur = 0.85
const usdToRub = 78.49

var eurToRub = usdToRub / usdToEur

func main() {
	getUserInput()
	fmt.Printf("Курс EUR к RUB: %.2f\n", eurToRub)
}

func getUserInput() {
	var currencyAmount float64
	var currencyFrom, currencyTo string

	fmt.Print("Введите количество валюты")
	fmt.Scan(&currencyAmount)
	fmt.Print("Введите из какой валюты нужно конвертировать?")
	fmt.Scan(&currencyFrom)
	fmt.Print("Введите в какую валюты нужно конвертировать?")
	fmt.Scan(&currencyTo)
}
