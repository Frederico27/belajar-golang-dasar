package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	//new address
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1)
	// fmt.Println(address2)

	// pointer everything to same
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
