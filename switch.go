package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, " + name + " Boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah pas")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length < 3:
		fmt.Println("Nama terlalu pendek")
	default:
		fmt.Println("Nama sudah pas")
	}

}
