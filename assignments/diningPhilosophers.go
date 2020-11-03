package main

import (
	"fmt"
	"sync"
	"time"
)

const NumMeals = 3
const NumChops = 5
const NumPhilos = 5
const NumHosts = 1
const HostSwitchDelayMs = 250

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS  *ChopS
	rightCS *ChopS
	id      int
}

type Host struct {
	whoCanEat map[int]bool
	mapLock   *sync.Mutex
	whoIsDone chan int
}

func (h *Host) canPhiloEat(philoId int) bool {
	h.mapLock.Lock()
	canPhiloEat := h.whoCanEat[philoId]
	h.mapLock.Unlock()
	return canPhiloEat
}

func (h *Host) runHost(wg *sync.WaitGroup) {
	defer wg.Done()
	reset := func() {
		for j := 0; j < NumPhilos; j++ {
			h.whoCanEat[j+1] = false
		}
	}
	fmt.Println("Hello philosophers, I am the host")
	i := 0
	countDone := 0
	for countDone < NumPhilos {
		h.mapLock.Lock()
		switch i {
		case 0:
			reset()
			h.whoCanEat[1] = true
			h.whoCanEat[3] = true
		case 1:
			reset()
			h.whoCanEat[2] = true
			h.whoCanEat[5] = true
		case 2:
			reset()
			h.whoCanEat[4] = true
		default:
			i = -1
		}
		h.mapLock.Unlock()

		select {
		case finishedEating := <-h.whoIsDone:
			fmt.Printf("Philosopher %d is done eating!\n", finishedEating)
			countDone++
		default:
			// no-op
		}
		i++
		time.Sleep(HostSwitchDelayMs * time.Millisecond)
	}
}

func (p Philo) eat(wg *sync.WaitGroup, host *Host) {
	signalDone := func() {
		host.whoIsDone <- p.id

	}
	defer signalDone()
	defer wg.Done()

	fmt.Println("Hello, I am philo:", p.id)
	i := 0
	for i < NumMeals {
		if host.canPhiloEat(p.id) {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("Philosopher %d is starting to eat meal %d \n", p.id, i)
			fmt.Printf("Philosopher %d is finishing eating meal %d \n", p.id, i)
			i++
			p.leftCS.Unlock()
			p.rightCS.Unlock()
		} else {
			// fmt.Println("Waiting to eat ", p.id)
		}
	}
}

func makePhilos() []*Philo {

	CSticks := make([]*ChopS, NumChops)
	for i := 0; i < NumChops; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, NumPhilos)
	for i := 0; i < NumPhilos; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}

	}
	return philos
}

func makeHost() *Host {
	whoCanEat := make(map[int]bool)
	for j := 0; j < NumPhilos; j++ {
		whoCanEat[j+1] = false
	}
	var mapLock sync.Mutex
	whoIsDone := make(chan int, NumPhilos)
	host := Host{whoCanEat, &mapLock, whoIsDone}
	return &host
}

func main() {
	philos := makePhilos()
	host := makeHost()

	var wg sync.WaitGroup
	wg.Add(NumPhilos + NumHosts)
	go host.runHost(&wg)
	for i := 0; i < NumPhilos; i++ {
		go philos[i].eat(&wg, host)
	}
	wg.Wait()
}
