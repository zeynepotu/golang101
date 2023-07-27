package main

import "fmt"

func main() {

	//Pointers
	/*var num = 10
	myPointer := &num
	fmt.Println(num)
	fmt.Println("Adresi:", myPointer)

	*myPointer = 25
	fmt.Println("Num yeni değer:", num)


	//Slices
	myArray := [5]int{1, 2, 3, 4, 5}

	slicedArray := myArray[0:2]

	fmt.Println(slicedArray)

	//Maps

	var plakalar = make(map[string]int)

	plakalar["İstanbul"] = 34
	plakalar["Ankara"] = 6
	plakalar["Çanakkale"] = 17
	plakalar["Bursa"] = 16
	fmt.Println("Plakalar", plakalar)
	delete(plakalar, "İstanbul")
	fmt.Println("Plakalar", plakalar)
	*/
	sayHello()
	sumNum(4, 5)

}

func sayHello() {
	fmt.Println("Selam")
}

func sumNum(num1, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}
