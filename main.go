package main

import (
	"fmt"
)

func main() {
	var number1, number2 float64
	var operatore string

	fmt.Print("Enter the first number :")
	fmt.Scanln(&number1)

	fmt.Print("Enter the operatore ( + - * /) :")
	fmt.Scanln(&operatore)

	fmt.Print("Enter the second number :")
	fmt.Scanln(&number2)

	switch operatore {

	case "+":
		fmt.Printf("%.2f %s %.2f = %.2f", number1, operatore, number2, number1+number2)

	case "-":
		fmt.Printf("%.2f %s %.2f = %.2f", number1, operatore, number2, number1-number2)

	case "*":
		fmt.Printf("%.2f %s %.2f = %.2f", number1, operatore, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("Divide by Zero situation")
		} else {
			fmt.Printf("%f %s %f = %f", number1, operatore, number2, number1/number2)

		}

	default:
		fmt.Println("Invalid Operator")

	}

}
