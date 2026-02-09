package dto

type BookInputdto struct {
	Title   string           `json:"title" binding:"required"`
	Genres  []string         `json:"genres" binding:"required,dive,checkgenre"`
	Authors []AuthorInputdto `json:"authors" binding:"required,dive,required"`
}

type AuthorBookResponsedto struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type BookResponsedto struct {
	ID      uint                    `json:"id"`
	Title   string                  `json:"title"`
	Genres  []string                `json:"genres"`
	Authors []AuthorBookResponsedto `json:"authors"`
}
