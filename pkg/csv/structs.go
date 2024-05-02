package csv

type Book struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publication-year"`
}
