package inc

import (
	"fmt"
	"math/rand"
	"sync"
)

type Account struct {
	Blnce int
	Mutex *sync.Mutex
}

var Accounts []*Account
var GlobalLock *sync.Mutex

func (a *Account) Widthdraw(v int) {
	// 데드락 회피방법1: 락의 범위를 좁게 잡는다.
	a.Mutex.Lock()
	a.Blnce -= v
	a.Mutex.Unlock()
}

func (a *Account) Deposit(v int) {
	// 데드락 회피방법1: 락의 범위를 좁게 잡는다.
	a.Mutex.Lock()
	a.Blnce += v
	a.Mutex.Unlock()
}

func (a *Account) Balance() int {
	blance := a.Blnce
	return blance
}

func Transfer(snder, recver int, money int) {
	/*
		[WRONG WAY]
		Accounts[snder].Mutex.Lock()
		fmt.Println("Lock:", snder)
		time.Sleep(1000 * time.Millisecond)
		Accounts[recver].Mutex.Lock()

		Accounts[snder].Widthdraw(money)
		Accounts[recver].Deposit(money)

		Accounts[recver].Mutex.Unlock()
		Accounts[snder].Mutex.Unlock()
	*/

	// 데드락 회피방법2: 락의 범위를 넓게 잡는다.
	GlobalLock.Lock()
	Accounts[snder].Widthdraw(money)
	Accounts[recver].Deposit(money)
	GlobalLock.Unlock()

	fmt.Println("Transfer => from:", snder, ", to:", recver, ", money:", money)
}

func GetTotalBalance() int {
	// 데드락 회피방법2: 락의 범위를 넓게 잡는다.
	GlobalLock.Lock()
	total := 0
	for i := 0; i < len(Accounts); i++ {
		total += Accounts[i].Balance()
	}
	GlobalLock.Unlock()

	return total
}

func RandomTransfer() {
	var snder, blnce int
	for {
		snder = rand.Intn(len(Accounts))
		blnce = Accounts[snder].Balance()
		if blnce > 0 {
			break
		}
	}

	var recver int
	for {
		recver = rand.Intn(len(Accounts))
		if snder != recver {
			break
		}
	}

	money := rand.Intn(blnce)
	Transfer(snder, recver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}
