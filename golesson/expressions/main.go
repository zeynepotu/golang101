package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	ahmet := Employee{
		Name:     "Ahmet Demir",
		Age:      27,
		Salary:   40000,
		FullTime: false,
	}

	ayşe := Employee{
		Name:     "Ayşe Kaya",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, ahmet)
	employees = append(employees, ayşe)

	for _, x := range employees {
		if x.Age >= 30 {
			fmt.Println(x.Name, "30 yaşında ya da daha büyük")
		} else {
			fmt.Println(x.Name, "30 yaşından küçük")
		}

		if x.Age > 30 && x.Salary > 50000 {
			fmt.Println(x.Name, "50000 den fazla maaş alıyor ve 30 yaşından büyük")
		} else {
			fmt.Println(x.Name, "50000 den az maaş alıyor veya 30 yaşından küçük")
		}

		if x.Age > 30 || x.Salary > 50000 {
			fmt.Println(x.Name, "50000 den fazla maaş alıyor veya 30 yaşından büyük")
		} else {
			fmt.Println(x.Name, "50000 den az maaş alıyor ve 30 yaşından küçük")
		}

		if (x.Age > 30 || x.Salary < 50000) && x.FullTime {
			fmt.Println(x.Name, "kriterlerimize uyuyor")
		} else {
			fmt.Println(x.Name, "kriterlerimize uymuyor")
		}

		fmt.Println("---------------------------------------------")
	}

}
