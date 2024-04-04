package main

import (
	"fmt"
	"its_task_2/wallet"
	"log"
)

func main() {
	wallet1 := wallet.NewWallet()
	wallet2 := wallet.NewWallet()

	err := wallet1.Deposit(1.23)
	if err != nil {
		log.Println(err)
	}
	err = wallet2.Deposit(10)
	if err != nil {
		log.Println(err)
	}
	err = wallet1.Withdraw(0.5)
	if err != nil {
		log.Println(err)
	}
	err = wallet2.Withdraw(15)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Wallet 1 balance:", wallet1.Balance())
	fmt.Println("Wallet 2 balance:", wallet2.Balance())
}
