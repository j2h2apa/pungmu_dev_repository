package main

import (
	"log"

	"github.com/go-playground/validator/v10"
)

//User :
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0, lte=130"`
	Email          string     `validate:"required, email"`
	FavouriteColor string     `validate:"iscolor"`
	Address        []*Address `validate:"required,dive,required"`
}

//Address :
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.validatem

func main() {
	validate = validator.New()

	validateVariable()
}

func validateVariable() {
	myEmail := "j2h2s2apa.gmail.com"

	errs := validate.Var(myEmail, "required,email")
	if errs != nil {
		log.Fatal(errs)
		return
	}
}