package csv

import (
	"encoding/csv"
	"os"
	"strconv"
)

func Append(path string, newBook *Book) {

	// Open csv file in append mode
	file, err := os.OpenFile(path, os.O_APPEND, 0644)
	if err != nil {
		panic("Error opening the csv file")
	}
	defer file.Close()

	//Creating a csv writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	year := strconv.Itoa(newBook.PublicationYear)
	//Converting newBook into slice of string
	newBookAdded := []string{newBook.Name, newBook.Author, year}

	//Writing to the csv file
	writer.Write([]string{})
	err = writer.Write(newBookAdded)
	if err != nil {
		panic(err)
	}
}
