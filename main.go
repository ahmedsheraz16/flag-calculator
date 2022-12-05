package main

import (
	"fmt"
)

func main() {
	var number1, number2 float32
	var flag string
	fmt.Print("Enter the first number:")
	fmt.Scanln(&number1)
	fmt.Print("Enter the second number:")
	fmt.Scanln(&number2)
	fmt.Print("Enter the operator (+ - * / %) : ")
	fmt.Scan(&flag)
	switch flag {
	case "+":
		fmt.Printf("%f %s %f = %f", number1, flag, number2, number1+number2)
	case "*":
		fmt.Printf("%f %s %f = %f", number1, flag, number2, number1*number2)
	case "-":
		fmt.Printf("%f %s %f = %f", number1, flag, number2, number1-number2)
	case "%":
		fmt.Printf("%f %s %f = %f", number1, flag, number2, number1+number2)
	case "/":
		if number2 == 0.0 {
			fmt.Println("Divide by zero situation")
		} else {
			fmt.Printf("%f %s %f = %f", number1, flag, number2, number1*number2)
		}
	default:
		fmt.Println("Invalid operator")
	}
}
