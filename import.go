package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.SayHello("Eko"))

	// fmt.Println(helper.version)    // tidak bisa diakses karena huruf awalan kecil
	// fmt.Println(helper.sayGoodBye) // tidak bisa diakses karena huruf awalan kecil
}
