package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearOfPublish int
}

func (b Book) toString() string {
	return fmt.Sprintf("Title:\t\t%q\n Author:\t\t%q\n YearOfPublish:\t\t%v\n", b.Title, b.Author, b.YearOfPublish)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "To Kill a MockingBird",
		Author:        "Harper Lee",
		YearOfPublish: 1960,
	},
	Book{
		ID:            2,
		Title:         "Rashmirathi",
		Author:        "Ramdhari Singh Dinkar",
		YearOfPublish: 1952,
	},
	Book{
		ID:            3,
		Title:         "A Wild Sheep Chase",
		Author:        "Haruki Murakami",
		YearOfPublish: 1982,
	},
	Book{
		ID:            4,
		Title:         "Moby Dick",
		Author:        "Herman Melville",
		YearOfPublish: 1851,
	},
	Book{
		ID:            5,
		Title:         "The Cuckoo's Calling",
		Author:        "Robert Galbraith",
		YearOfPublish: 2013,
	},
	Book{
		ID:            6,
		Title:         "The Old Man and the Sea",
		Author:        "Ernest Hemingway",
		YearOfPublish: 1952,
	},
	Book{
		ID:            7,
		Title:         "Animal farm",
		Author:        "George Orwell",
		YearOfPublish: 1945,
	},
	Book{
		ID:            8,
		Title:         "To Have and Have Not",
		Author:        "Ernest Hemingway",
		YearOfPublish: 1937,
	},
	Book{
		ID:            9,
		Title:         "Three Men in a Boat",
		Author:        "Jerome K. Jerome",
		YearOfPublish: 1889,
	},
	Book{
		ID:            10,
		Title:         "The Alchemist",
		Author:        "Paulo Coelho",
		YearOfPublish: 1960,
	},
	Book{
		ID:            11,
		Title:         "The Catcher in the Rye",
		Author:        "J. D. Salinger",
		YearOfPublish: 1951,
	},
	Book{
		ID:            12,
		Title:         "Ulysses",
		Author:        "James Joyce",
		YearOfPublish: 1922,
	},
}
