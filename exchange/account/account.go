package account

import (
	"errors"
	"fmt"
	"sync"
)

func NewAccount(Id, Name string) *Account {
	return &Account{
		Id:      Id,
		Name:    Name,
		Balance: 0,
	}
}

type Account struct {
	Id      string
	Name    string
	Balance float64
	Lock    sync.Mutex
}

func (a *Account) SaveMoney(amount float64) {
	a.Balance = a.Balance + amount
}

func (a *Account) DrawMoney(amount float64,index int) error {
	//a.Lock.Lock()
	//defer a.Lock.Unlock()
	cur := a.Balance
	if cur < amount {
		return errors.New("余额不足。")
	}
	fmt.Println(fmt.Sprintf("第%d个人取款%d元。", index, 5000))
	//time.Sleep(1 * time.Second)
	a.Balance = cur - amount
	return nil
}

func (a *Account) QueryBalance() float64 {
	return a.Balance
}
