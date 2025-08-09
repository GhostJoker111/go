package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operationType, operationTypeErr := scanOperationType()
	if operationTypeErr != nil {
		fmt.Println(operationTypeErr)
		return
	}

	numbers, err := scanNumbers()
	if err != nil {
		fmt.Printf("Ошибка преобразования: %v\n", err)
		return
	}

	sort.Float64s(numbers)
	operationResult := calculateOperation(operationType, numbers)
	fmt.Printf("Результат операции %s - %.2f\n", operationType, operationResult)

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

	if operationType == "AVG" {
		for _, value := range numbers {
			result += value
		}
		result = result / float64(len(numbers))
	} else if operationType == "SUM" {
		for _, value := range numbers {
			result += value
		}
	} else if operationType == "MED" {
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

func scanNumbers() ([]float64, error) {
	var input string
	var numbers []float64

	fmt.Println("Введите число или числа через запятую (n для выхода): ")
	fmt.Scan(&input)

	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, fmt.Errorf("Ошибка преобразования '%s': %v", part, err)
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}
