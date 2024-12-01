package inventory

import (
	"bufio"
	"fmt"
	filestorage "inventory-store-cli/filestorage"
	"inventory-store-cli/models"
	"inventory-store-cli/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

var STORE_FILENAME = "inventory-store.txt"
var MainOptions = []string{
	"add a item",
	"remove a item",
	"list the inventory items",
	"exit the menu",
}
var RemoveOptions = []string{
	"by passing unique id",
	"by going through each one and then decide",
	"exit the menu",
}

type Inventory struct {
	Store []models.InventoryItem
	Name  string
}

func (i *Inventory) LoadItems() {
	items, err := filestorage.ConvertFileItemsToStructs(STORE_FILENAME)
	fmt.Println(err)
	fmt.Println("struct items are")
	fmt.Println(len(items))

	fmt.Println(items)
	i.Store = items
}

func (i *Inventory) Add() {
	item := ReadInventory()
	fileStore := &filestorage.FileStorage{Filename: STORE_FILENAME}
	itemString := fmt.Sprintf("%s->%v->%v->%s->%v", item.Name, item.Price, item.ManufacturedDate, item.Brand, item.Id)
	fileStore.Write(itemString)
	i.Store = append(i.Store, item)
	fmt.Printf("\nAdded item %+v", item)
}

func (i *Inventory) Remove() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		RenderMenu(RemoveOptions)
		fmt.Println("please select how you want to remove the item")
		scanner.Scan()
		selectedOption := scanner.Text()
		switch selectedOption {
		case "0":
			fmt.Printf("you have selected to remove it by entering uniqueId")
			fmt.Printf("\nplease enter the unique id\n")
			scanner.Scan()
			itemToBeRemovedId := scanner.Text()

			parsedItemToBeRemovedId, err := strconv.ParseInt(itemToBeRemovedId, 10, 64)
			utils.ErrorHandler(err)
			itemIndex := -1
			for i, item := range i.Store {
				if item.Id == parsedItemToBeRemovedId {
					itemIndex = i
					break
				}
			}
			if itemIndex == -1 {
				fmt.Printf("for passed unique id, item doesn't exist")
				fmt.Printf("please check and try again")
			} else {
				i.Store = append(i.Store[:itemIndex], i.Store[itemIndex+1:]...)
				go filestorage.RemoveItemFromFile(STORE_FILENAME, parsedItemToBeRemovedId)
				fmt.Printf("item with %v id removed successfully", itemToBeRemovedId)
			}
		case "1":
			fmt.Printf("\ngoing through each item and deciding it")

			for index, item := range i.Store {

				fmt.Printf("\n%+v", item)
				fmt.Printf("\ndo you want to remove this ?, if yes, press y or n for checking the next item")
				scanner.Scan()
				wantToRemove := scanner.Text()
				if strings.ToLower(wantToRemove) == "y" || strings.ToLower(wantToRemove) == "yes" {
					fmt.Printf("\nyou have decided to remove this item")
					go filestorage.RemoveItemFromFile(STORE_FILENAME, i.Store[index].Id)
					i.Store = append(i.Store[:index], i.Store[index+1:]...)
					fmt.Printf("\nremoved successfully")
					return
				}

			}

			return
		case "2":
			return
		}

	}

}

func (i Inventory) ListAllItems() {
	if len(i.Store) == 0 {
		fmt.Printf("Inventory store is empty")
		return
	}
	for _, item := range i.Store {
		fmt.Printf("%+v\n", item)
	}
}

func RenderMenu(menu []string) {
	fmt.Printf("\n")
	for i, m := range menu {
		fmt.Printf("%v - %v\n", i, m)
	}

}

func ReadInventory() models.InventoryItem {
	scanner := bufio.NewScanner(os.Stdin)

	var item models.InventoryItem
	fmt.Println("please enter the item name")
	scanner.Scan()
	item.Name = scanner.Text()
	fmt.Println("please enter the price")
	scanner.Scan()
	inputVal := scanner.Text()
	price, err := strconv.ParseFloat(inputVal, 64)
	utils.ErrorHandler(err)
	item.Price = price
	timelLayout := "2006-01-02 15:04:05"
	fmt.Printf("\nenter the manufactured date (format %v), if you omit it, current time will be considerd", timelLayout)
	var manufTime string
	scanner.Scan()
	manufTime = scanner.Text()
	if manufTime == "" && len(manufTime) == 0 {
		currTime := time.Now()
		formattedTime := currTime.Format(timelLayout)
		item.ManufacturedDate = utils.ConvertStringToTime(formattedTime)

	} else {
		item.ManufacturedDate = utils.ConvertStringToTime(manufTime)
	}

	fmt.Println("please enter the brand name")
	scanner.Scan()
	item.Brand = scanner.Text()

	item.Id = time.Now().UnixNano()

	return item

}

func CreateStore() Inventory {
	store := Inventory{}
	store.LoadItems()
	fmt.Println(store)

	return store
}

func init() {
	fmt.Printf("init called")
}

type ReaderWriter interface {
	Write(data string) error
	Read(data string) error
}

func WriteItem(rw ReaderWriter, item string) {

	err := rw.Write(item)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
