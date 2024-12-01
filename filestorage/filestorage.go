package filestorage

import (
	"bufio"
	"fmt"
	"inventory-store-cli/models"
	"inventory-store-cli/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

type FileStorage struct {
	Filename string
}

func (fw *FileStorage) Write(data string) error {
	file, err := os.OpenFile(fw.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = file.WriteString(data + "\n")

	if err != nil {
		return err
	}

	return nil
}

func (fr *FileStorage) Read() (string, error) {
	data, err := os.ReadFile(fr.Filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ConvertFileItemsToStructs(filename string) ([]models.InventoryItem, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	defer file.Close()
	var items []models.InventoryItem

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var item models.InventoryItem
	for scanner.Scan() {
		line := scanner.Text()
		splittedLine := strings.Split(line, "->")
		fmt.Println(splittedLine)

		item.Name = splittedLine[0]
		price, err := strconv.ParseFloat(splittedLine[1], 64)
		if err != nil {
			log.Printf("failed to convert price to float when reading from the file, error is %v", err)
		}
		item.Price = price
		item.ManufacturedDate = utils.ConvertStringToTime(splittedLine[2])
		item.Brand = splittedLine[3]
		id, err := strconv.ParseInt(splittedLine[4], 10, 64)
		if err != nil {
			log.Printf("failed to convert string id to int64 when reading from the file,error is %v", err)
		}
		item.Id = id
		items = append(items, item)
	}

	fmt.Println(items)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func RemoveItemFromFile(filename string, itemId int64) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		item := scanner.Text()
		splitItemString := strings.Split(item, "->")

		itemIdString := strconv.FormatInt(itemId, 10)
		if itemIdString != splitItemString[4] {
			lines = append(lines, item)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	file, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	writer.Flush()
	defer file.Close()
	return nil
}
