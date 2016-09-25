// Package account solves an Exercism challenge.
package account

import (
	"sync"
)

// Account holds all the data of a bank account.
type Account struct {
	b int64
	c bool
	sync.Mutex
}

// Open creates a new Account with initial balance init.
func Open(init int64) *Account {
	if init < 0 {
		return nil
	}
	return &Account{b: init}
}

// Close closes Account a.
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.c {
		return 0, false
	}
	a.c = true
	return a.b, true
}

// Balance gets the current balance of Account a.
func (a *Account) Balance() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.c {
		return -1, false
	}
	return a.b, true
}

// Deposit alters the amount of money in Account a by d.
func (a *Account) Deposit(d int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.c {
		return -1, false
	}
	if a.b+d < 0 {
		return -1, false
	}
	a.b += d
	return a.b, true
}
