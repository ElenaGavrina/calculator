package main

import(
	"fmt"

)

func main(){
	var num1, num2 int
	var op string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)
	fmt.Print("Enter the operator ( + - * / ): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		fmt.Printf("%d %s %d = %d",num1,op,num2,num1 + num2)
	case "-":
		fmt.Printf("%d %s %d = %d",num1,op,num2,num1 - num2)
	case "*":
		fmt.Printf("%d %s %d = %d",num1,op,num2,num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Divide by Zero situation")
		}else{
			fmt.Printf("%d %s %d = %d",num1,op,num2,num1 / num2)
		}
	default:
		fmt.Println("Invalid operator")
	}

}