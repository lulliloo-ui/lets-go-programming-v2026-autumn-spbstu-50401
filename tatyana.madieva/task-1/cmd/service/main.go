package main

import "fmt"

func main() {
	var s_v1 string
	fmt.Scan(&s_v1)
	var value1 int
	_, err1 := fmt.Sscanf(s_v1, "%d", &value1)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	var s_v2 string
	fmt.Scan(&s_v2)
	var value2 int
	_, err2 := fmt.Sscanf(s_v2, "%d", &value2)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	var s_op string
	fmt.Scan(&s_op)
	var res int
	switch s_op {
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
