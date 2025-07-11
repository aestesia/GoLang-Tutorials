package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	//ownerInfo owner
	//owner
}

// type owner struct {
// 	name string
// }

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func main() {
	//var myEngine gasEngine = gasEngine{25, 15, owner{"John"}}
	var myEngine gasEngine = gasEngine{25, 15}
	//myEngine.mpg = 15
	//fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	//fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

	// Anonymous struct (cannot be reuseable)
	// var myEngine2 = struct{
	// 	mpg uint8
	// 	gallons uint8
	// }{25,15}

	//fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Total miles left in a tank %v", myEngine.milesLeft())

}
