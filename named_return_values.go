package main

import "fmt"

func getCompleteName() (firstname, middlename, lastname string) {
	firstname = "Eko"
	middlename = "Kurniawan"
	lastname = "Khannedy"

	return firstname, middlename, lastname
}

func main() {

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
