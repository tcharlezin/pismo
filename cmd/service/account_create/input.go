package account_create

type AccountCreateInput struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

func (input *AccountCreateInput) Validate() bool {
	return true
}
