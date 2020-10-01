package main

import (
	"fmt"

	"github.com/mannerism/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("juno")
	account.Deposit(10)
	fmt.Println(account)
}
