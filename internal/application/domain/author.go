package domain

type AuthorInput struct {
	Username string
}

type GetAuthorQuery struct {
	Username *string 
	ID *uint
}
