package main

import (
	"sync"
	"time"

	acnt "lec.youtube.tucker/thread.usage/inc"
)

func main() {
	for i := 0; i < 20; i++ {
		acnt.Accounts = append(acnt.Accounts, &acnt.Account{Blnce: 1000, Mutex: &sync.Mutex{}})
	}
	acnt.GlobalLock = &sync.Mutex{}

	acnt.PrintTotalBalance()

	for i := 0; i < 10; i++ {
		go acnt.GoTransfer()
	}

	for {
		acnt.PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}
