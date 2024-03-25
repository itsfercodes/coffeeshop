package main

import (
	"coffeeshop/coffee"
	"fmt"
)

func main() {
	fmt.Println("Printing the list of coffees available")
	coffees, err := coffee.GetCoffees()

	if err != nil {
		fmt.Println("Error while getting coffeelist ", err)
		return
	}

	// Prints all coffees available
	for _, element := range coffees.List {
		fmt.Printf("%s for $%v\n", element.Name, element.Price)
	}

	fmt.Println("Is Latte Available? ", coffee.IsCoffeeAvailable("Latte"))
}
