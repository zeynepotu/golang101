package main

import "fmt"

type Car struct {
	model    string
	year     int
	door_num int
}

func main() {

	//var myArray [4]int

	myArray2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(myArray2[0])

	myCar := Car{
		model:    "BMW",
		year:     2015,
		door_num: 4,
	}

	fmt.Printf("Arabamın modeli %s yılı %d kapı sayısı da %d dür.", myCar.model, myCar.year, myCar.door_num)

}
