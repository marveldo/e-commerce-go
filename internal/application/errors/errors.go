package apperrors

import (
	"errors"

	"gorm.io/gorm"
)

func CheckDuplicatekeyError(err error) ( bool ){
	if err == nil {
		return false 
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true 
	}else {
       return false 
	}
}

func UpdateNotFoundError(err error) (bool){
	return  err.Error() == "Query Object Not Found"
}
