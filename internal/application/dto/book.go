package dto

type BookInputdto struct {
	Title       string           `json:"title" binding:"required"`
	Genres      []string         `json:"genres" binding:"required,dive,checkgenre"`
	Authors     []AuthorInputdto `json:"authors" binding:"required,dive,required"`
	Price       int64            `json:"price" binding:"required"`
	Description string           `json:"description" binding:"required"`
}

type AuthorBookResponsedto struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type BookResponsedto struct {
	ID          uint                    `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Genres      []string                `json:"genres"`
	Price       int64                   `json:"price"`
	Authors     []AuthorBookResponsedto `json:"authors"`
}

type BookQueryDto struct {
	Title string  `form:"title" binding:"omitempty,min=1"`
	Genres []string `form:"genre" binding:"omitempty,min=1"` 
}
