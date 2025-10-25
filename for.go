package main

import "fmt"

func main() {

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke -", counter)
	// 	counter++
	// }

	// fmt.Println("Perulangan Selesai")

	// init statement dan post statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke -", counter)
	}

	fmt.Println("Perulangan Selesai")

	names := []string{"Eko", "Kurniawan", "Khannedy"}

	for i := 0; i < len(names); i++ {
		fmt.Println("Index", i, "=", names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

}
