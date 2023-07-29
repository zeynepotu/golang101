package main

import (
	"fmt"
)

func main() {
	// 1'den 10'a kadar olan sayıları ekrana yazdıralım (for döngüsü)
	fmt.Println("For Döngüsü ile 1'den 10'a kadar olan sayılar:")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 1'den 10'a kadar olan sayıları ekrana yazdıralım (while döngüsü gibi)
	fmt.Println("While Döngüsü ile 1'den 10'a kadar olan sayılar:")
	j := 1
	for j <= 10 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()

	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println("Dizideki sayılar:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
