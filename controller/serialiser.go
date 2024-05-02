package controller

import "library/pkg/csv"

type AllBooksRequest struct{}
type AllBooksResponse struct {
	AdminBooks   []csv.Book `json:"admin-books"`
	RegularBooks []csv.Book `json:"regular-books"`
}

type AddBookRequest struct {
	Name            string `json:"name" binding:"required"`
	Author          string `json:"author" binding:"required"`
	PublicationYear int    `json:"publication-year" binding:"required"`
}

type AddBookResponse struct {
	Message string `json:"msg"`
	Error   error  `json:"error"`
}

type DeleteBookRequest struct {
	Name string `json:"name" binding:"required"`
}
type DeleteBookResponse struct {
	Message string `json:"msg"`
	Error   error  `json:"error"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Message string `json:"msg"`
	Token   string `json:"token"`
	Error   error  `json:"error"`
}
