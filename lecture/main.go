package main

import (
  "fmt"
  "log"

  "github.com/ralpioxxcs/go_study/lecture/accounts"
)

func main() {
  account := accounts.NewAccount("ralpio")
  account.Deposit(500)
  fmt.Println(account)

  err := account.Withdraw(600)
  if err != nil {
    log.Fatalln(err)
  }
  fmt.Println(account)
}

