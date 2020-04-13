package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*Account : 통장*/
type Account struct {
	balance int // 잔액
	mutex   *sync.Mutex
}

/*Widthdraw : 출금*/
func (a *Account) Widthdraw(val int) {
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

/*Deposit : 입금*/
func (a *Account) Deposit(val int) {
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

/*Balance : 통장잔액 return*/
func (a *Account) Balance() int {
	a.mutex.Lock()
	balance := a.balance
	a.mutex.Unlock()
	return balance
}

/*통장 리스트 slice*/
var accounts []*Account
var globalLock *sync.Mutex

/*Transfer : entity 에 입출금 동시 작업*/
func Transfer(sender, receiver int, money int) {
	globalLock.Lock()
	accounts[sender].Widthdraw(money)
	accounts[receiver].Deposit(money)
	globalLock.Unlock()
}

/*GetTotalBalance : 통장 slice 잔액의 합*/
func GetTotalBalance() int {
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	return total
}

/*RandomTransfer : 랜덤하게 입출금 호출*/
func RandomTransfer() {
	var sender, balance int

	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int

	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

/*GoTransfer : 무한 호출*/
func GoTransfer() {
	for {
		RandomTransfer()
	}
}

/*PrintTotalBalance : 전체 입금액 출력*/
func PrintTotalBalance() {
	globalLock.Lock()
	fmt.Printf("Total : %d\n", GetTotalBalance())
	globalLock.Unlock()
}
