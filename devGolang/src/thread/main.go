package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// case 1
	// go func1()
	// for i := 0; i < 20; i++ {
	// 	time.Sleep(100 * time.Millisecond)
	// 	fmt.Println("main :", i)
	// }
	// fmt.Scanln()
	// case 1 end

	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
	}

	globalLock = &sync.Mutex{}

	for i := 0; i < 1; i++ {
		go GoTransfer()
	}

	for {
		PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}

func func1() {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("func1 :", i)
	}
}
