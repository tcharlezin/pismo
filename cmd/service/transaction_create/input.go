package transaction_create

type TransactionCreateInput struct {
	AccountID       uint    `json:"account_id"`
	OperationTypeID uint    `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func (input *TransactionCreateInput) Validate() bool {
	if input.OperationTypeID < 1 || input.OperationTypeID > 4 {
		return false
	}

	if input.Amount < 0 {
		return false
	}

	return true
}
