package main

import "fmt"

func main() {
	type NoKTP string

	var ktpEko NoKTP = "111111"
	fmt.Println(ktpEko)
	var contoh = "222222"
	var contohNoKTP NoKTP = NoKTP(contoh)
	fmt.Println(contohNoKTP)
}
