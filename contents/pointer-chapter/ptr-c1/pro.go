package ptrc1

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(cust *customer, trans transaction) error {
	if trans.transactionType == transactionDeposit {
		cust.balance += trans.amount
		return nil
	} else if trans.transactionType == transactionWithdrawal {
		if cust.balance < trans.amount {
			return errors.New("insufficient funds")
		}
		cust.balance -= trans.amount
		return nil
	} else {
		return errors.New("unknown transaction type")
	}
}
