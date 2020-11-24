package accounts 

import "errors"

var errNoMoney = errors.New("Can't withdraw you are poor!")

// Account struct
type Account struct {
  owner   string
  balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account{
  account := Account{owner: owner, balance: 0}
  return &account
}

// Deposit x amount on your account
// don't copy of account ,, use address
func (a *Account) Deposit(amount int) {
  a.balance += amount
}

// Balance return balance on your account
func (a Account) Balance() int {
  return a.balance
}

// Withdraw x amount from on your account
func (a *Account) Withdraw(amount int) error {
  if a.balance < amount {
    //return errors.New("Can'y withdraw you are poor!")
    return errNoMoney
  }
  a.balance -= amount
  return nil
}
