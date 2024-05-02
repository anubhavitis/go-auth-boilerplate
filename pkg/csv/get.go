package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func Get(path string) []Book {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	allBooks := []Book{}

	// Print each record
	for i, row := range records {
		if i == 0 {
			continue
		}
		year, _ := strconv.Atoi(row[2])

		newBook := Book{
			Name:            row[0],
			Author:          row[1],
			PublicationYear: year,
		}
		allBooks = append(allBooks, newBook)
	}

	return allBooks
}
