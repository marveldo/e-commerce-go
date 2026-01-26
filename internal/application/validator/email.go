package validator

import (
	"net/mail"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Emailformat validator.Func = func (fl validator.FieldLevel) bool {
   email := fl.Field().Interface().(string)
   err , _ := mail.ParseAddress(email)
   if err != nil {
	return true
   }
   return false
}


var PasswirdFormat validator.Func = func (fl validator.FieldLevel) bool {
   password := fl.Field().Interface().(string)
   is_upper := regexp.MustCompile(`[A-Z]`).MatchString(password)
   is_special := regexp.MustCompile(`[!@#~$%^&*]`).MatchString((password))
   return is_upper && is_special
}
