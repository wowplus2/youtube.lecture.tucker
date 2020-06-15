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
	a.Mutex.Lock()
	a.Blnce -= v
	a.Mutex.Unlock()
}

func (a *Account) Deposit(v int) {
	a.Mutex.Lock()
	a.Blnce += v
	a.Mutex.Unlock()
}

func (a *Account) Balance() int {
	a.Mutex.Lock()
	blance := a.Blnce
	a.Mutex.Unlock()

	return blance
}

func Transfer(snder, recver int, money int) {
	GlobalLock.Lock()
	Accounts[snder].Widthdraw(money)
	Accounts[recver].Deposit(money)
	GlobalLock.Unlock()
}

func GetTotalBalance() int {
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
