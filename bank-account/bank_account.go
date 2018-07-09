package account

type Bank struct {
	balance int
	open    bool
	ch      *chan bool
}

func Open(sum int) *Bank {
	if sum < 0 {
		return nil
	}
	ch := make(chan bool, 1)
	bank := Bank{balance: sum, open: true, ch: &ch}
	bank.Unlock()

	return &bank
}

// lock bank using chan
func (b *Bank) Lock() {
	<-*b.ch
}

// unlock bank using chan
func (b *Bank) Unlock() {
	*b.ch <- true
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
