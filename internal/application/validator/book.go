package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/marveldo/gogin/internal/db"
)

var BookGenreValidator validator.Func = func(fl validator.FieldLevel) bool {
	genre := fl.Field().Interface().(string)
	allowedGenres := []db.BookGenre{
		db.ART,
		db.CHILDREN,
		db.NonFiction,
		db.ScienceFiction,
		db.Fantasy,
		db.COMICS,
		db.HEALTH,
		db.Horror,
		db.Mystery,
		db.RELIGION,
		db.Romance,
		db.SELF_HELP,
		db.THRILLER,
		db.TRAVEL,
		db.ACTION,
		db.ADVENTURE,
}

	if genre != "" && genre!= " " && containsGenre(genre, allowedGenres) {
		return true
	}else {
		return false
	}

	
}

func containsGenre(genre string, allowedGenres []db.BookGenre) bool {
	for _, g := range allowedGenres {
		if strings.EqualFold(string(g), genre) {
			return true
		}
	}
	return false
}