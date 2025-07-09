package main

import "fmt"

const usdToEur = 0.85
const usdToRub = 78.49

var eurToRub = usdToRub / usdToEur

func main() {
	fmt.Printf("Курс EUR к RUB: %.2f\n", eurToRub)
}
