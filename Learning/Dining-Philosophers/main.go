package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name string
	leftFork int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Krishau", leftFork: 4, rightFork: 0},
	{name: "Sumit", leftFork: 0, rightFork: 1},
	{name: "Yash", leftFork: 1, rightFork: 2},
	{name: "Rounak", leftFork: 2, rightFork: 3},
	{name: "Dipayan", leftFork: 3, rightFork: 4},
}

var hunger = 2; // how many times a philosopher will eat
var eatTime = 1*time.Second; // How long it takes to eat
var thinkTime = 3*time.Second; // How long it takes to think
var sleepTime = 1*time.Second; // How long it takes to sleep

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("----------------------------")
	fmt.Println("The Table is empty.")
	// start the meal
	dine();
}

func dine() {
	eatTime = 0*time.Second;
	thinkTime = 0*time.Second;
	sleepTime = 0*time.Second;
	// wg is the waitgroup that keeps track of how many philosophers are still at the table. When 
	// it reaches zero, everyone is finished eating and has left. We add 5 (the number of philosophers) to
	// this waitgroup
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	//we want everyone to be seated before they start eating, so create a waitgroup for that, and set it to 5.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	//forks is a map of all 5 forks. Forks are assigned using the fields leftfork and rightfork in the philosopher type.
	//Each fork, then, can be found using the index(an integer), and each fork has a unique mutex.
	forks := make(map[int]*sync.Mutex)
	for i:=range philosophers {
		forks[i] = &sync.Mutex{}
	}
	// start the meal by iterating  through our slice of philosophers
	for i:=range philosophers {
		go dinningProblem(philosophers[i], forks, wg, seated)
	}
	// wait for the philosophers to finish. This blocks until the waitgroup is 0
	wg.Wait();
}

func dinningProblem(philosopher Philosopher, forks map[int]*sync.Mutex, wg *sync.WaitGroup, seated *sync.WaitGroup) {
	defer wg.Done()
	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name);
	//Decrement the seated waitgroup by one
	seated.Done()
	// wait untill everyone is seated
	seated.Wait()
	// start eating
	for range hunger {
		fmt.Printf("%s is hungry.\n", philosopher.name)
		// Example: Change order for even philosophers to left first
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s has the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s has the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s has the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s has the right fork.\n", philosopher.name)
		}

		// eat
		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)
		// THe philosopher starts to think, but does not put down the forks
		fmt.Printf("%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		// put down the left fork
		forks[philosopher.leftFork].Unlock()
		// put down the right fork
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("%s put down both the fork.\n", philosopher.name)

		fmt.Println(philosopher.name, "left the table.")
	}
}