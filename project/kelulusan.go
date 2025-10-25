package main

import "fmt"

func main() {

	nilaiSiswa := map[string]int{
		"Joko":  10,
		"Anwar": 5,
	}

	for nama, nilai := range nilaiSiswa {
		switch {
		case nilai >= 6:
			fmt.Printf("Selamat %s anda lulus dengan nilai %d\n", nama, nilai)
		case nilai <= 5:
			fmt.Printf("Mohon Maff %s anda tidak lulus dengan nilai %d\n", nama, nilai)
		}
	}
}
