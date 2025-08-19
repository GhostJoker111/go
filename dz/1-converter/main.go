package main

import (
	"fmt"
	"strings"
)

var allCurrencies = []string{"USD", "EUR", "RUB"}
var exchangeRates = map[string]float64{
	"USD": 1.0,
	"EUR": 1.15,
	"RUB": 0.0125,
}

func main() {
	currencyAmount, currencyFrom, currencyTo := getUserInput()
	convertCurrency(currencyAmount, currencyFrom, currencyTo)
}

func getUserInput() (float64, string, string) {
	var currencyAmount float64
	var currencyFrom, currencyTo string

	for {
		fmt.Printf("Введите из какой валюты нужно конвертировать(%s)?", strings.Join(allCurrencies, ", "))
		fmt.Scan(&currencyFrom)
		currencyFrom = strings.ToUpper(currencyFrom)

		if currencyFrom == "USD" || currencyFrom == "EUR" || currencyFrom == "RUB" {
			break
		}
		fmt.Println("Введены некорректные символы. Выберите предложенную валюту")
	}

	for {
		fmt.Println("Введите количество валюты: ")
		currAmount, err := fmt.Scan(&currencyAmount)
		if err != nil {
			fmt.Println("Неверно указано количество. Введите число")
			continue
		}

		if currAmount < 0 {
			fmt.Println("Введите положительное число")
			continue
		}
		break
	}

	for {
		availableCurrencies := getAvailableCurrencies(currencyFrom)
		fmt.Printf("Введите в какую валюты нужно конвертировать?(%s)", strings.Join(availableCurrencies, ", "))
		fmt.Scan(&currencyTo)
		currencyTo = strings.ToUpper(currencyTo)

		if currencyTo != "USD" && currencyTo != "EUR" && currencyTo != "RUB" {
			fmt.Println("Введены некорректные символы. Выберите предложенную валюту")
			continue
		}

		if currencyTo == currencyFrom {
			fmt.Println("Выбрана одинаковая валюта для конвертации. Выберите другую валюту")
			continue
		}

		break
	}

	return currencyAmount, currencyFrom, currencyTo
}

func getAvailableCurrencies(currencyFrom string) []string {
	var availableCurrencies []string

	for _, currency := range allCurrencies {
		if currency != currencyFrom {
			availableCurrencies = append(availableCurrencies, currency)
		}
	}

	return availableCurrencies
}

func convertCurrency(amount float64, currencyFrom, currencyTo string) float64 {
	fromRate, fromExists := exchangeRates[currencyFrom]
	toRate, toExists := exchangeRates[currencyTo]

	if !fromExists || !toExists {
		fmt.Println("Ошибка: валюта не найдена в таблице курсов")
		return 0
	}

	amountInUSD := amount * fromRate
	result := amountInUSD / toRate
	fmt.Println(result)

	return result
}
