package main

import "fmt"

type Account struct {
	Name    string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.Balance += amount
		fmt.Printf("%s have been deposit $%.2f. Now the balance are $%.2f\n", a.Name, amount, a.Balance)
	}
}

func (a *Account) WithDraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid amount")
	}

	if amount > a.Balance {
		fmt.Println("The balance insuficient")
	}

	a.Balance -= amount
}

func Transfer(from, to *Account, amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid amount transfer")
	}

	if from.Balance < amount {
		fmt.Println("The amount insuficient")
	}

	from.Balance -= amount
	to.Balance += amount

	fmt.Printf("Success transfered amount $%2.f from %s to %s\n", amount, from.Name, to.Name)
}
