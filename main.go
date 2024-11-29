package main

import (
	"bufio"
	"fmt"
	"inventory-store-cli/inventory"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("started the inventory store manager")

	store := inventory.CreateStore()

	var input string

	for {
		inventory.RenderMenu(inventory.MainOptions)
		fmt.Println("please provide the input and select the options")
		scanner.Scan()

		input = scanner.Text()
		fmt.Println("you have entered value", input)

		switch input {
		case "0":
			fmt.Println("add the items")
			store.Add()
		case "1":
			fmt.Println("remove the item")
			store.Remove()
		case "2":
			fmt.Println("List the inventory items")
			store.ListAllItems()
		case "3":
			fmt.Println("exiting the store manager")
			os.Exit(0)
		}
	}
}
