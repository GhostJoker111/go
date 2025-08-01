package main

import (
	"fmt"
	"strings"
)

var allCurrencies = []string{"USD", "EUR", "RUB"}

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
		currencyAmount, err := fmt.Scan(&currencyAmount)
		if err != nil {
			fmt.Println("Неверно указано количество. Введите число")
			continue
		}

		if currencyAmount < 0 {
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

func convertCurrency(currencyAmount float64, currencyFrom, currencyTo string) float64 {
	var usdToRub, eurToRub, usdToEur = 0.0125, 0.011, 1.15
	var convertCurrency float64

	switch {
	case currencyFrom == "USD" && currencyTo == "RUB":
		convertCurrency = currencyAmount / usdToRub
		fmt.Printf("%f", convertCurrency)
	case currencyFrom == "RUB" && currencyTo == "USD":
		convertCurrency = currencyAmount * usdToRub
		fmt.Printf("%f", convertCurrency)
	case currencyFrom == "EUR" && currencyTo == "RUB":
		convertCurrency = currencyAmount / eurToRub
		fmt.Printf("%f", convertCurrency)
	case currencyFrom == "RUB" && currencyTo == "EUR":
		convertCurrency = currencyAmount * eurToRub
		fmt.Printf("%f", convertCurrency)
	case currencyFrom == "USD" && currencyTo == "EUR":
		convertCurrency = currencyAmount / usdToEur
		fmt.Printf("%f", convertCurrency)
	case currencyFrom == "EUR" && currencyTo == "USD":
		convertCurrency = currencyAmount * usdToEur
		fmt.Printf("%f", convertCurrency)

	default:
		fmt.Print("Что-то пошло не так")
	}

	return convertCurrency
}
