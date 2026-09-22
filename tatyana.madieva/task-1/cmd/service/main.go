package main

import "fmt"

func main() {
	var value1 int
	_, err1 := fmt.Scan(&value1)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var value2 int
	_, err2 := fmt.Scan(&value2)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var sOp string
	_, err3 := fmt.Scan(&sOp)
	if err3 != nil {
		fmt.Println("Invalid operation")
		return
	}
	var res int
	switch sOp {
	case "+":
		res = value1 + value2
	case "-":
		res = value1 - value2
	case "*":
		res = value1 * value2
	case "/":
		if value2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		res = value1 / value2
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Println(res)
}
