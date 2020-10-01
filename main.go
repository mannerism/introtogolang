package main

import (
	"fmt"

	"github.com/mannerism/learngo/mydict"
)

func main() {
	// Account
	// account := accounts.NewAccount("juno")
	// account.Deposit(10)
	// fmt.Println(account)

	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}

	hello, err := dictionary.Search(word)
	fmt.Println(hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
