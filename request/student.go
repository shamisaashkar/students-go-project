package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Student struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (req Student) Validate() error {

	return validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.Required, validation.Length(1, 255), is.UTFLetterNumeric),
		validation.Field(&req.LastName, validation.Required, validation.Length(1, 255), is.UTFLetterNumeric),
		validation.Field(&req.ID, validation.Required, validation.Max(402243012), validation.Min(90999999)),
	)

}
