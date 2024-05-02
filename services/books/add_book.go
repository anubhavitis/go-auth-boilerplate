package books

import (
	"errors"
	"library/constants"
	"library/pkg/csv"
)

func AddBook(newBook *csv.Book) error {

	// validate if the newBook exists
	for _, book := range csv.Get(constants.RegularUserCSV) {
		if book.Name == newBook.Name {
			return errors.New("book already exists")
		}
	}
	csv.Append(constants.RegularUserCSV, newBook)
	return nil
}
