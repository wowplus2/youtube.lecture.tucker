package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

func MakeTire(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Tire,"

		outChan <- car
	}
}

func MakeEngine(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Engine, "

		outChan <- car
	}
}

func StartWork(ch chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		ch <- Car{val: "Car" + strconv.Itoa(i) + ": "}
		i++
	}
}

func main() {
	ch01 := make(chan Car)
	ch02 := make(chan Car)
	ch03 := make(chan Car)

	go StartWork(ch01)
	go MakeTire(ch01, ch02)
	go MakeEngine(ch02, ch03)

	for {
		result := <-ch03
		fmt.Println(result.val)
	}
}
