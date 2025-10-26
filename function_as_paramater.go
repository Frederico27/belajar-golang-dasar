package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filtered_name := filter(name)

	fmt.Println("Hello", filtered_name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

}
