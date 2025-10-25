package main

import "fmt"

func main() {
	name := "Jokofadfasd"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Printf("Hi, %s Boleh kenalan?", name)
	}

	//if short name
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	}
}
