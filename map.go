package main

import "fmt"

func main() {

	// var person map[string]string = map[string]string{}

	// person["name"] = "Eko"
	// person["address"] = "Indonesia"
	// person["language"] = "Golang"

	person := map[string]string{
		"name":     "Eko",
		"address":  "Indonesia",
		"language": "Golang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["language"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Eko Kurniawan Khannedy"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")

	fmt.Println(book)

}
