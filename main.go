package main

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/oleksandrkhmil/code-generation-poc/form3"
)

func Pointer[T any](v T) *T { return &v }

func main() {
	account := form3.Account{
		Record: form3.Record{
			ID:             Pointer(strfmt.UUID4("123")),
			OrganisationID: Pointer(strfmt.UUID4("234")),
			Type:           form3.RecordTypePayment,
		},
		Attributes: &form3.AccountAttributes{
			AccountNumber: "12345678",
		},
	}
	if err := account.Validate(strfmt.NewFormats()); err != nil {
		fmt.Println(err.Error())
	}
}
