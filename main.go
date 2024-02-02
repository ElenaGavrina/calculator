package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) (c int) {
	return a / b
}

func main() {
	var a int
	var b int
	var c string
	fmt.Println("Введите первое число:")
	fmt.Scanln(&a)
	fmt.Println("Введите второе число:")
	fmt.Scanln(&b)
	fmt.Println("Введите действие: + для сложения, - для вычитания, * для умножения, / для деления")
	fmt.Scanln(&c)
	switch{
	case c == "+":
		fmt.Println(add(a, b))
	case c == "-":
		fmt.Println(sub(a, b))
	case c == "*":
		fmt.Println(mult(a,b))
	case c == "/":
		if b == 0{
			fmt.Println("Делить на ноль нельзя!")
			return
		}
		fmt.Println(div(a,b))
	}

}
