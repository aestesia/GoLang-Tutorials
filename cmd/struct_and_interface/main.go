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
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("made it")
	} else {
		fmt.Println("fuel up")
	}
}

func main() {
	//var myEngine gasEngine = gasEngine{25, 15, owner{"John"}}

	var myEngine gasEngine = gasEngine{25, 15}
	//var myEngine electricEngine = electricEngine{25, 15}

	//myEngine.mpg = 15
	//fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	//fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

	// Anonymous struct (cannot be reuseable)
	// var myEngine2 = struct{
	// 	mpg uint8
	// 	gallons uint8
	// }{25,15}

	//fmt.Println(myEngine.mpg, myEngine.gallons)
	canMakeIt(myEngine, 50)

	fmt.Printf("Total miles left in a tank %v", myEngine.milesLeft())

}
