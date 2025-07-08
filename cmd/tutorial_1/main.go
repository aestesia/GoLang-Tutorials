package main

import (
	"errors"
	"fmt"
)

func main() {

	// fmt.Println("Hello World")
	// var intNum int = 32767
	// fmt.Println(intNum)

	// var floatNum float64 = 12345678.9
	// fmt.Println(floatNum)

	// var floatNum32 float32 = 10.1
	// var intNum32 int32 = 2
	// var result = floatNum32 + float32(intNum32)
	// fmt.Println(result)

	// var intNum1 int = 3
	// var intNum2 int = 2
	// fmt.Println(intNum1 / intNum2)
	// fmt.Println(intNum1 % intNum2)

	// var myString string = "Hello World"
	// fmt.Println(myString)

	// fmt.Println(utf8.RuneCountInString("A"))

	// var myRune rune = 'a'
	// fmt.Println(myRune)

	// var myBoolean bool = false
	// fmt.Println(myBoolean)

	// //var myVar = "text"
	// myVar := "text"
	// fmt.Println(myVar)

	// var1, var2 := 1, 2
	// fmt.Println(var1, var2)

	// const myConst string = "can't be changed"
	// fmt.Println(myConst)

	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Println("the result of division is %v", result)
	// } else {
	// 	fmt.Println("The result of division is %v with remainder %v", result, remainder)
	// }

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("the result of division is %v", result)
	default:
		fmt.Printf("The result of division is %v with remainder %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by 0")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
