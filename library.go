package library

import "fmt"

type Book struct {
	Id    int
	Title string
}

type IShelf interface {
	Search(id int) *Book
}

type shelf struct {
	Books []*Book
}

func (s *shelf) Search(id int) *Book {
	for _, book := range s.Books {
		if book.Id == id {
			return book
		}
	}
	fmt.Println("book not found")
	return nil
}

func NewShelf() IShelf {
	return &shelf{
		Books: []*Book{
			{
				Id:    1,
				Title: "Math",
			},
			{
				Id:    2,
				Title: "Physics",
			},
			{
				Id:    3,
				Title: "Computer",
			},
		},
	}
}
