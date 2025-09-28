package main

import "fmt"

func main() {

	name := "Eko Kurniawan Khannedy"
	fmt.Println(name)

	name = "Eko"
	fmt.Println(name)

	var (
		firstName  = "Eko"
		middleName = "Kurniawan"
		lastName   = "Khannedy"
		umur       = 30
	)

	fmt.Println(firstName + " " + middleName + " " + lastName + " umurku " + fmt.Sprint(umur) + " tahun")

}
