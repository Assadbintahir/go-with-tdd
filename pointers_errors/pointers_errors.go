package pointers_errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

// errors pkg by dave cheney is used to add context & stacktrace to errors. (compatible with stdlib)
var ErrInsufficientFunds = errors.New("can't withdraw due to insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
