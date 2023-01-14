package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Go Calculator")
	cmd := readLine("Enter command: [a]dd, [s]ubtract,[m]ultiply,[d]ivide:")
	fmt.Println(cmd)
	if cmd == "a" {
		num1, num2 := getUserNumber()
		result := num1 + num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "s" {
		num1, num2 := getUserNumber()
		result := num1 - num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "m" {
		num1, num2 := getUserNumber()
		result := num1 * num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "d" {
		num1, num2 := getUserNumber()
		result := float32(num1) / float32(num2)
		sResult := fmt.Sprintf("%0.2f", result)
		fmt.Println(sResult)
	} else {
		fmt.Println("Invalid input")
	}
}

func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumber() (int, int) {
	num1String := readLine("First Number: ")
	num1, err := strconv.Atoi(num1String)
	if err != nil {
		fmt.Println(err)
	}
	num2String := readLine("Second Number: ")
	num2, err := strconv.Atoi(num2String)
	if err != nil {
		fmt.Println(err)
	}
	return num1, num2
}
