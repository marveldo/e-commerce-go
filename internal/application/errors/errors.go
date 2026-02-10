package apperrors

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckDuplicatekeyError(err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	} else {
		return false
	}
}
func CheckNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	} else {
		return false
	}
}

func UpdateNotFoundError(err error) bool {
	q := err.Error() == "Query Object Not Found"
	q1 := errors.Is(err, gorm.ErrRecordNotFound)
	return q || q1
}

func ForeignKeyConstraintError(err error) bool {
	return errors.Is(err, gorm.ErrForeignKeyViolated)
}

func PasswordIncorrect(err error) bool {
	return err.Error() == "Password Not Correct"
}

func InvalidTokenError(err error) bool {
	msg := err.Error()
	return strings.Contains(msg , "idtoken:")
	}


func ErrorFormat(g *gin.Context, err error) {
	if CheckDuplicatekeyError(err) {
		g.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})

	} else if UpdateNotFoundError(err) {
		g.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  err.Error(),
		})

	} else if CheckNotFoundError(err) {
		g.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  err.Error(),
		})

	} else if PasswordIncorrect(err) {
		g.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err.Error(),
		})

	} else if ForeignKeyConstraintError(err) {
		g.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
	} else if InvalidTokenError(err) {
		g.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err.Error(),
		})
	} else {
		g.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})

	}
}
