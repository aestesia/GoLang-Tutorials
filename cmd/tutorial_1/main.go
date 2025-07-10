package main

import (
	"fmt"
	"strings"
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

	// -------------------------------------------------------------------
	// var printValue string = "Hello World"
	// printMe(printValue)

	// var numerator int = 11
	// var denominator int = 2
	// var result, remainder, err = intDivision(numerator, denominator)

	// // if err != nil {
	// // 	fmt.Println(err.Error())
	// // } else if remainder == 0 {
	// // 	fmt.Println("the result of division is %v", result)
	// // } else {
	// // 	fmt.Println("The result of division is %v with remainder %v", result, remainder)
	// // }

	// switch {
	// case err != nil:
	// 	fmt.Println(err.Error())
	// case remainder == 0:
	// 	fmt.Printf("the result of division is %v", result)
	// default:
	// 	fmt.Printf("The result of division is %v with remainder %v", result, remainder)
	// }

	// switch remainder {
	// case 0:
	// 	fmt.Printf("The division was exact")
	// case 1, 2:
	// 	fmt.Printf("The division was close")
	// default:
	// 	fmt.Printf("The division was not close")
	// }

	// -------------------------------------------------------------------

	// //var intArr [3]int32 = [3]int32{1,2,3}
	// //intArr := [3]int32{1, 2, 3}
	// intArr := [...]int32{1, 2, 3}
	// // intArr[1] = 123
	// // fmt.Println(intArr[0])
	// // fmt.Println(intArr[1:3])
	// fmt.Println(intArr)

	// var intSlice []int32 = []int32{4, 5, 6}
	// fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	// intSlice = append(intSlice, 7)
	// fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))

	// var intSlice2 []int32 = []int32{8, 9}
	// intSlice2 = append(intSlice, intSlice2...)
	// fmt.Println(intSlice)

	// var intSlice3 []int32 = make([]int32, 3, 10)
	// fmt.Println(intSlice3)

	// var myMap map[string]uint8 = make(map[string]uint8)
	// fmt.Println(myMap)

	// var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	// fmt.Println(myMap2["Adam"])
	// var age, ok = myMap2["Jason"]
	// if ok {
	// 	fmt.Println("The age is %v", age)
	// } else {
	// 	fmt.Println("invalid name")
	// }

	// for name, age := range myMap2 {
	// 	fmt.Printf("Name: %v Age: %v\n", name, age)
	// }

	// for i, v := range intArr {
	// 	fmt.Printf("Index: %v, Value: %v \n", i, v)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// -------------------------------------------------------------------

	//var myString = "résumé"
	var myString = []rune("résumé")
	var indexed = myString[0]
	fmt.Printf("%v, %T", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"t", "e", "s", "t", "i", "n", "g"}
	var strBuilder strings.Builder
	for i := range strSlice {
		//catStr += strSlice[i]
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}

// func printMe(printValue string) {
// 	fmt.Println(printValue)
// }

// func intDivision(numerator int, denominator int) (int, int, error) {
// 	var err error
// 	if denominator == 0 {
// 		err = errors.New("Cannot divide by 0")
// 		return 0, 0, err
// 	}
// 	var result int = numerator / denominator
// 	var remainder int = numerator % denominator
// 	return result, remainder, err
// }
