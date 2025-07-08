package main

import "fmt"

func main() {
	const usdToEur = 0.85
	const usdToRub = 78.49
	eurToRub := usdToRub / usdToEur

	fmt.Printf("Курс EUR к RUB: %.2f\n", eurToRub)

}
