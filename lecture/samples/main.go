package main

import (
	"fmt"
	"log"

	"github.com/ralpioxxcs/go_study/lecture/accounts"
	"github.com/ralpioxxcs/go_study/lecture/dict"
)

func main() {
	account := accounts.NewAccount("ralpio")
	account.Deposit(500)
	fmt.Println(account)

	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Balance(), account.Owner())

	//////////////////////////////////////////////////
	fmt.Println("\nDictionary")

	dictionary := dict.Dictionary{"first": "First word"}

	word := "hello"
	definition := "greeting"
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}

	hello, _ := dictionary.Search(word)
	fmt.Println(hello)

	err3 := dictionary.Add(word, definition)
	if err3 != nil {
		fmt.Println(err3)
	}

	err4 := dictionary.Update("asdasd", "Second")
	if err4 != nil {
		fmt.Println(err4)
	}
	updateword, _ := dictionary.Search(word)
	fmt.Println(updateword)

	//dictionary.Delete(word)
	dictionary.Delete("wow")

}
