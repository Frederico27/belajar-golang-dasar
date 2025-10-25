package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Khannedy"
}

func main() {

	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)

	firstname2, _ := getFullName()
	fmt.Println(firstname2)

}
