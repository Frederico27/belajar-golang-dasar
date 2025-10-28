package main

import (
	"fmt"
)

func main() {

	alice := Account{Name: "Alice", Balance: 2000}
	bob := Account{Name: "Bob", Balance: 1000}

	alice.Deposit(100)
	bob.WithDraw(500)

	fmt.Println(alice.Balance)
	fmt.Println(bob.Balance)

	Transfer(&alice, &bob, 100)

	fmt.Println(alice.Balance)
	fmt.Println(bob.Balance)

}
