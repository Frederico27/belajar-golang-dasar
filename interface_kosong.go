package main

import "fmt"

func Ups() any {

	//semua tipe data
	// return 1
	// return true
	return "Hello"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
