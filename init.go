package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" //taru blank di depan agar hanya mengeksekusi init()
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
