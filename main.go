package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello to my calculator!")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in your first number: ")
	scanner.Scan()
	nOne, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println("Type in your second number: ")
	scanner.Scan()
	nTwo, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println("Choose your math operation:")
	fmt.Println("1 addition")
	fmt.Println("2 subtraction")
	fmt.Println("3 multiplication")
	fmt.Println("4 division")
	fmt.Println("Type in the number in front of the math operation: ")
	scanner.Scan()
	mathSign, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	switch mathSign {
	case 1:
		var result int64 = nOne + nTwo
		fmt.Println("Your Result is: " + strconv.FormatInt(result, 16))
	case 2:
		var result int64 = nOne - nTwo
		fmt.Println("Your Result is: " + strconv.FormatInt(result, 16))
	case 3:
		var result int64 = nOne * nTwo
		fmt.Println("Your Result is: " + strconv.FormatInt(result, 16))
	case 4:
		var result int64 = nOne / nTwo
		fmt.Println("Your Result is: " + strconv.FormatInt(result, 16))
	}
}
