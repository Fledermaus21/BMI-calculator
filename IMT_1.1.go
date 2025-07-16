package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	for {
		height, weight := getUserInput()
		IMT, err := calculateIMT(height, weight)
		if err != nil {
			// fmt.Println(err)
			// continue
			fmt.Println()
			panic(err)
		}
		fmt.Printf("Ваш индекс массы тела %.1f\n", IMT)
		fmt.Println(identifyCondition(IMT))
		if getUserChoice() {
			continue
		} else {
			break
		}
	}
}

func calculateIMT(height, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("Не указан вес или рост")
	}
	result := weight / math.Pow(height/100, 2)
	return result, nil
}

func getUserInput() (float64, float64) {
	var kg, cm float64
	fmt.Print("Введите свой рост в см: ")
	fmt.Scan(&cm)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&kg)
	return cm, kg
}

func identifyCondition(coef float64) string {
	var result string
	switch {
	case coef <= 16:
		result = "Выраженный дефицит массы тела"
	case coef <= 18.5:
		result = "Недостаточная масса тела"
	case coef <= 25:
		result = "Норма"
	case coef <= 30:
		result = "Избыточная масса тела"
	case coef <= 35:
		result = "Ожирение первой степени"
	case coef <= 40:
		result = "Ожирение второй степени"
	default:
		result = "Ожирение третьей степени"
	}
	return result
}

func getUserChoice() bool {
	fmt.Println("\n Хотите повторить расчет? [Д/Н]: ")
	var answer string
	fmt.Scan(&answer)
	if answer == "Д" || answer == "д" {
		return true
	} else {
		return false
	}
}
