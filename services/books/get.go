package books

import (
	"library/constants"
	"library/pkg/csv"
)

func GetAllBooks() ([]csv.Book, []csv.Book) {
	regularBooks := csv.Get(constants.RegularUserCSV)
	adminBooks := csv.Get(constants.AdminUserCSV)

	// create new slice of both books that only contains names

	return regularBooks, adminBooks
}
