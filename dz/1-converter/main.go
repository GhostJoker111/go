package main

import "fmt"

const usdToEur = 0.85
const usdToRub = 78.49

func main() {
	eurToRub := usdToRub / usdToEur

	fmt.Printf("Курс EUR к RUB: %.2f\n", eurToRub)

}
