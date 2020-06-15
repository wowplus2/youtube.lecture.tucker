package main

import (
	"sync"
	"time"

	acnt "lec.youtube.tucker/thread.deadlock/inc"
)

func main() {
	for i := 0; i < 20; i++ {
		// 데드락 회피방법1: 락의 범위를 좁게 잡는다.
		acnt.Accounts = append(acnt.Accounts, &acnt.Account{Blnce: 1000, Mutex: &sync.Mutex{}})
	}

	// 데드락 회피방법2: 락의 범위를 넓게 잡는다.
	acnt.GlobalLock = &sync.Mutex{}
	go func() {
		for {
			acnt.Transfer(0, 1, 100)
		}
	}()

	go func() {
		for {
			acnt.Transfer(1, 0, 100)
		}
	}()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
