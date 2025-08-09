package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func main() {
	var numbers []float64
	operationType, operationTypeErr := scanOperationType()

	if operationTypeErr != nil {
		fmt.Println(operationTypeErr)
		return
	}

	for {
		number := scanNumbers()
		if number == 0 {
			break
		}
		numbers = append(numbers, number)
	}
	sort.Float64s(numbers)
	operationResult := calculateOperation(operationType, numbers)
	fmt.Printf("Результат операции %.0s - %.0f", operationType, operationResult)

}

func scanOperationType() (string, error) {
	var operationType string
	AVG, SUM, MED := "AVG", "SUM", "MED"

	fmt.Println("Введите тип операции: ")
	fmt.Scan(&operationType)
	operationType = strings.ToUpper(operationType)

	if operationType == AVG {
		return AVG, nil
	} else if operationType == SUM {
		return SUM, nil
	} else if operationType == MED {
		return MED, nil
	} else {
		return "", errors.New("Неизвестная операция. Повторите опытку")
	}
}

func calculateOperation(operationType string, numbers []float64) float64 {
	var result float64

	for _, value := range numbers {
		if operationType == "AVG" {
			result += value / float64(len(numbers))
			continue
		} else if operationType == "SUM" {
			result += value
			continue
		}

		if len(numbers)%2 == 1 {
			result = numbers[len(numbers)/2]
		} else {
			mid1 := numbers[len(numbers)/2-1]
			mid2 := numbers[len(numbers)/2]
			result = (mid1 + mid2) / 2
		}
	}
	return result
}

func scanNumbers() float64 {
	var number float64
	fmt.Println("Введите число (n для выхода): ")
	fmt.Scan(&number)

	return number
}
