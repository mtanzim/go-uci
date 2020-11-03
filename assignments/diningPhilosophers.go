package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS  *ChopS
	rightCS *ChopS
	id      int
}

const NumMeals = 3
const NumChops = 5
const NumPhilos = 5

func (p Philo) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	// NOTE: make this an infinite loop to see the deadlock
	for i := 0; i < NumMeals; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("Starting to eat ", p.id)
		fmt.Println("Finishing eating", p.id)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}

func initPhilos() ([]*ChopS, []*Philo) {

	CSticks := make([]*ChopS, NumChops)
	for i := 0; i < NumChops; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, NumPhilos)
	for i := 0; i < NumPhilos; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}

	return CSticks, philos
}

func main() {
	fmt.Println("Hello philos")
	_, philos := initPhilos()
	for _, p := range philos {
		fmt.Println("Hello, I am philo:", p.id)
	}
	var wg sync.WaitGroup

	wg.Add(NumPhilos)
	for i := 0; i < NumPhilos; i++ {
		go philos[i].eat(&wg)
	}
	wg.Wait()

}
