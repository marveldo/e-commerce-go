package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ValidatorType struct {
	key string 
	validator validator.Func
}



func RegisterValidator(o ValidatorType) {
	v , ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation(o.key , o.validator)
	}
}

func AddValidators() []ValidatorType {
	validators := []ValidatorType{
	   {"checkemail", Emailformat},
	   {"checkpassword", PasswirdFormat},
	}
	return validators
}

func RegisterAllValidators(){
	validators := AddValidators()
	for _ , v := range validators {
		RegisterValidator(v)
	}
}