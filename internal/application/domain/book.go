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
}

type BookInputWithoutAuthors struct {
	Title  string
	Genres []string
}

type GetBookQuery struct {
	ID *uint
	Title *string
	Genres []*string
}

