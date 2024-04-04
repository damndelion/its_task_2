package wallet

import (
	"fmt"
	"sync"
)

type Wallet struct {
	bitcoin float64
	mutex   sync.Mutex
}

func NewWallet() *Wallet {
	return &Wallet{bitcoin: 0, mutex: sync.Mutex{}}
}

func (w *Wallet) Deposit(amount float64) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount < 0 {
		return fmt.Errorf("deposit amount cannot be negative")
	}
	w.bitcoin += amount

	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if amount < 0 {
		return fmt.Errorf("withdraw amount cannot be negative")
	}

	if amount > w.bitcoin {
		return fmt.Errorf("not enough funds: requested %v, available %v", amount, w.bitcoin)
	}
	w.bitcoin -= amount

	return nil
}

func (w *Wallet) Balance() float64 {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.bitcoin
}
