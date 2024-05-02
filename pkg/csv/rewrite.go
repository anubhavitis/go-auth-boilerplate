package csv

import (
	"encoding/csv"
	"os"
	"strconv"
)

func Rewrite(path string, payload []Book) {

	// Open the CSV file again for writing
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the filtered Books back to the CSV file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	book_string := [][]string{
		{
			"Book Name", "Author", "Publication Year",
		},
	}

	for _, book := range payload {
		year := strconv.Itoa(book.PublicationYear)
		book_string = append(book_string, []string{book.Name, book.Author, year})
	}

	for _, book := range book_string {
		err := writer.Write(book)
		if err != nil {
			panic(err)
		}
	}
}
