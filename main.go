package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	Name       string
	Eating     bool
	LeftHand   *Fork
	RightHand  *Fork
	TimesEated int
}

type Fork struct {
	Ready     bool
	Available sync.Mutex
}

var philosophers []*Philosopher
var wg sync.WaitGroup

func main() {

	fmt.Println("Philosophers are sitting in the table...")

	forkOne := Fork{Ready: true, Available: sync.Mutex{}}
	forkTwo := Fork{Ready: true, Available: sync.Mutex{}}
	forkThree := Fork{Ready: true, Available: sync.Mutex{}}
	forkFour := Fork{Ready: true, Available: sync.Mutex{}}
	forkFive := Fork{Ready: true, Available: sync.Mutex{}}

	philosophers = []*Philosopher{
		{Name: "Plato", Eating: false, RightHand: &forkFive, LeftHand: &forkOne},
		{Name: "Socrates", Eating: false, RightHand: &forkOne, LeftHand: &forkTwo},
		{Name: "Aristotle", Eating: false, RightHand: &forkTwo, LeftHand: &forkThree},
		{Name: "Pascal", Eating: false, RightHand: &forkThree, LeftHand: &forkFour},
		{Name: "Locke", Eating: false, RightHand: &forkFour, LeftHand: &forkFive},
	}

	wg.Add(len(philosophers))

	// spawn one goroutine for each philosopher
	for _, philosopher := range philosophers {
		// spawn a goroutine
		go DiningProblem(philosopher)
	}

	wg.Wait()
	fmt.Println("DONE!")
}

func DiningProblem(philosopher *Philosopher) {
	defer wg.Done()

	// Try to eat
	for {
		if philosopher.TimesEated >= 3 {
			break
		}

		fmt.Println(philosopher.Name, " is hungry... ")

		// If are available, lock the forks
		philosopher.LeftHand.Available.Lock()
		philosopher.RightHand.Available.Lock()

		philosopher.Eating = true
		philosopher.LeftHand.Ready = false
		philosopher.RightHand.Ready = false

		// Eat
		fmt.Println("\t", philosopher.Name, " is eating... ")
		time.Sleep(3 * time.Second)

		philosopher.RightHand.Ready = true
		philosopher.LeftHand.Ready = true
		philosopher.TimesEated++
		philosopher.Eating = false

		// Unlock
		philosopher.LeftHand.Available.Unlock()
		philosopher.RightHand.Available.Unlock()
		time.Sleep(time.Second / 2)
	}
}
