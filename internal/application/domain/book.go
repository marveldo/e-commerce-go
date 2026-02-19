package domain

// type BookGenre struct {
// 	Name string
// }

// type Author struct {
// 	Username string
// }

type BookInput struct {
	Title   string
	Genres  []string
	Authors []AuthorInput
	Price   int64
	Description string
}

type BookInputWithoutAuthors struct {
	Title  string
	Price   int64
	Description string
	Genres []string
}

type GetBookQuery struct {
	ID uint
	Title string
	Genres []string
	Price   int64
	Description string
}

