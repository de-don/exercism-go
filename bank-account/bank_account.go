package account

import "sync"

type Bank struct {
	balance int
	open    bool
	sync.Mutex
}

func Open(sum int) *Bank {
	if sum < 0 {
		return nil
	}
	bank := Bank{balance: sum, open: true}

	return &bank
}

func (b *Bank) Balance() (int, bool) {
	if b.open {
		return b.balance, true
	}
	return 0, false
}

func (b *Bank) Close() (int, bool) {
	b.Lock()
	defer b.Unlock()

	if !b.open {
		return 0, false
	}
	b.open = false
	return b.balance, true
}

func (b *Bank) Deposit(sum int) (int, bool) {
	b.Lock()
	defer b.Unlock()

	if (!b.open) || (b.balance+sum < 0) {
		return 0, false
	}
	b.balance += sum
	return b.balance, true
}
