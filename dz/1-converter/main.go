package main

import "fmt"

func main() {
	currencyAmount, currencyFrom, currencyTo := getUserInput()
	convertCurrency(currencyAmount, currencyFrom, currencyTo)
}

func getUserInput() (float64, string, string) {
	var currencyAmount float64
	var currencyFrom, currencyTo string

	fmt.Print("Введите количество валюты")
	fmt.Scan(&currencyAmount)
	fmt.Print("Введите из какой валюты нужно конвертировать?")
	fmt.Scan(&currencyFrom)
	fmt.Print("Введите в какую валюты нужно конвертировать?")
	fmt.Scan(&currencyTo)

	return currencyAmount, currencyFrom, currencyTo
}

func convertCurrency(currencyAmount float64, currencyFrom, currencyTo string) {}
