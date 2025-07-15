package main

// type contactInfo struct {
// 	Name  string
// 	Email string
// }

// type purchaseInfo struct {
// 	Name   string
// 	Price  float32
// 	amount int
// }

//struct generics
type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	//Slice
	// var intSlice = []int{1, 2, 3}
	// fmt.Println(sumSlice[int](intSlice))

	// var float32Slice = []float32{1, 2, 3}
	// fmt.Println(sumSlice[float32](float32Slice))

	//Any situation
	// var intSlice = []int{}
	// fmt.Println(isEmpty(intSlice))

	// var float32Slice = []float32{1, 2, 3}
	// fmt.Println(isEmpty(float32Slice))

	//when can't infer the type of generic parameter (io/ioutil is depracated so can't run)
	// var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	// fmt.Printf("\n%+v", contacts)

	// var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	// fmt.Printf("\n%+v", purchases)

	//struct generics
	// var gasCar = car[gasEngine]{
	// 	carMake:  "Honda",
	// 	carModel: "Civic",
	// 	engine: gasEngine{
	// 		gallons: 12.4,
	// 		mpg:     40,
	// 	},
	// }

	// var electriCar = car[electricEngine]{
	// 	carMake:  "Tesla",
	// 	carModel: "Model 3",
	// 	engine: electricEngine{
	// 		kwh:   57.5,
	// 		mpkwh: 4.17,
	// 	},
	// }

}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slices []T) bool {
	return len(slices) == 0
}

// func loadJSON[T contactInfo | purchaseInfo](filepath string) T {
// 	data, _ = ioutil.ReadFile(filepath)

// 	var loaded = []T{}
// 	json.Unmarshal(data, &loaded)

// 	return loaded
// }
