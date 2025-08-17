package main

import "fmt"

var Reset = "\033[0m"
var Green = "\033[32m"

func printMenu() {
	fmt.Println(Green + " __| |_____________________________________________________________________| |__" + Reset)
	fmt.Println(Green + " __   _____________________________________________________________________   __" + Reset)
	fmt.Println(Green + "   | |                                                                     | |" + Reset)
	fmt.Println(Green + "   | |   ____             ____      _            _       _                 | |" + Reset)
	fmt.Println(Green + "   | |  / ___| ___       / ___|__ _| | ___ _   _| | __ _| |_ ___  _ __     | |" + Reset)
	fmt.Println(Green + "   | | | |  _ / _ \\     | |   / _` | |/ __| | | | |/ _` | __/ _ \\| '__|    | |" + Reset)
	fmt.Println(Green + "   | | | |_| | (_) |    | |__| (_| | | (__| |_| | | (_| | || (_) | |       | |" + Reset)
	fmt.Println(Green + "   | |  \\____|\\___/      \\____\\__,_|_|\\___|\\__,_|_|\\__,_|\\__\\___/|_|     _ | |" + Reset)
	fmt.Println(Green + "   | | | __ ) _   _ _     / \\  | (_) |  \\/  | ___   ___   __ _ _____   _(_)| |" + Reset)
	fmt.Println(Green + "   | | |  _ \\| | | (_)   / _ \\ | | | | |\\/| |/ _ \\ / _ \\ / _` / __\\ \\ / / || |" + Reset)
	fmt.Println(Green + "   | | | |_) | |_| |_   / ___ \\| | | | |  | | (_) | (_) | (_| \\__ \\\\ V /| || |" + Reset)
	fmt.Println(Green + "   | | |____/ \\__, (_) /_/   \\_\\_|_| |_|  |_|\\___/ \\___/ \\__,_|___/ \\_/ |_|| |" + Reset)
	fmt.Println(Green + "   | |        |___/                                                        | |" + Reset)
	fmt.Println(Green + " __| |_____________________________________________________________________| |__" + Reset)
	fmt.Println(Green + " __   _____________________________________________________________________   __" + Reset)
	fmt.Println(Green + "   | |                                                                     | |" + Reset)
	fmt.Println("Enter number of operation: ")
	fmt.Println("1. Sum")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("0. Exit")
}

func getInput() (float64, float64) {
	var num1, num2 float64
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func main() {
	for {
		fmt.Printf("\x1bc")
		printMenu()

		var operation string
		fmt.Scanln(&operation)

		if operation == "0" {
			break
		}

		var num1, num2, result float64

		switch operation {
		case "1":
			num1, num2 = getInput()
			result = num1 + num2
			fmt.Println("Result: ", result)

		case "2":
			num1, num2 = getInput()
			result = num1 - num2
			fmt.Println("Result: ", result)

		case "3":
			num1, num2 = getInput()
			result = num1 * num2
			fmt.Println("Result: ", result)

		case "4":
			num1, num2 = getInput()
			if num2 != 0 {
				result = num1 / num2
				fmt.Println("Result: ", result)
			} else {
				fmt.Println("Error: Division by zero is not allowed.")
			}

		default:
			fmt.Println("Invalid input, please try again.")
		}

		fmt.Println("========== Press a key to continue ==========")
		fmt.Scanln()
	}
}
