package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{90, 80, 70}

	fmt.Println(values)

	var values2 = [...]int{
		90,
		80,
		70,
		100,
	}

	fmt.Println(len(values2))
	fmt.Println(values2[0])
	fmt.Println(values2[1])

}
