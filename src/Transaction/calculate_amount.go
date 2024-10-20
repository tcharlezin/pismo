package Transaction

import (
	"pismo/cmd/handle/transaction_create"
)

const PURCHASE = 1
const INSTALLMENT_PURCHASE = 2
const WITHDRAWAL = 3
const PAYMENT = 4

func CalculateAmount(input transaction_create.TransactionCreateInput) float64 {
	if input.OperationTypeID == PAYMENT {
		return input.Amount
	}

	return input.Amount * (-1)
}
