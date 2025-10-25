package main

import "fmt"

func tambah(a int, b int) int {
	return a + b
}

func kurang(a int, b int) int {
	return a - b
}

func kali(a int, b int) int {
	return a * b
}

func main() {

	var a, b int
	var c string

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)
	fmt.Print("Pilih operasi (+, -, *): ")
	fmt.Scanln(&c)

	switch c {
	case "+":
		fmt.Printf("Hasil: %d\n", tambah(a, b))
	case "-":
		fmt.Printf("Hasil: %d\n", kurang(a, b))
	case "*":
		fmt.Printf("Hasil: %d\n", kali(a, b))
	default:
		fmt.Println("Operasi tidak dikenal")
	}

}
