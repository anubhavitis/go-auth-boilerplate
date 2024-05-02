package books

import (
	"errors"
	"library/constants"
	"library/pkg/csv"
)

func DeleteBook(bookName string) error {

	// validate if the book is present or not

	allBooks := csv.Get(constants.RegularUserCSV)

	newBooks := []csv.Book{}
	bookFound := false
	for _, book := range allBooks {
		if book.Name != bookName {
			newBooks = append(newBooks, book)
		} else {
			bookFound = true
		}
	}

	if !bookFound {
		return errors.New("book not found")
	}

	csv.Rewrite(constants.RegularUserCSV, newBooks)
	return nil
}
