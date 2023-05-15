package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	
	"github.com/mnogu/go-calculator"
)

func main() {
	fmt.Println("Enter the data to calculate: ")	
	
	reader := bufio.NewReader(os.Stdin)
	val, err := reader.ReadString('\n')
	if err != nil {
    	    fmt.Println(err)
  	}

	result, err:= calculator.Calculate(val)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
