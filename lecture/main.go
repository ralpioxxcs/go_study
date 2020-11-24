package main

import (
  "fmt"

  "github.com/ralpioxxcs/go_study/lecture/accounts"
)

func main() {
  account := accounts.NewAccount("ralpio")
  fmt.Println(account)
}

