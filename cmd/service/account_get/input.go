package account_get

type AccountGetInput struct {
	AccountID uint `json:"account_id" validate:"required"`
}

func (input *AccountGetInput) Validate() bool {

	return true
}
