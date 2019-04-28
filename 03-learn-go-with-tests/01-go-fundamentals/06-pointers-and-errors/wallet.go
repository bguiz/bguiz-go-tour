package wallet

import (
	"errors"
	"fmt"
)

type Ether float32

var ErrInsufficientFunds = errors.New("insufficient balance")

// NOTE
// - create the error at the package level, instead of in line where it
//   is thrown, because that way it can be referenced from within tests too

func (e Ether) String() string {
	return fmt.Sprintf("%.2f ETH", e)
}

// NOTE
// - think of this as a type alias

type Wallet struct {
	balance Ether
}

func (w *Wallet) Deposit(amount Ether) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Ether) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Ether {
	return w.balance
}

// NOTE
// - cannot use `func (w Wallet)` because it creates a copy of the wallet
// - instead, use a pointer, which is a reference to the original wallet object
//   using `func (w *Wallet)`
// - the callsite *does not* need to change its call
