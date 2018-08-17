package main

import (
"fmt"
"time"
"sync"
)

var forks [5]sync.Mutex

func main() {

	for i := range forks {
		go philo(i,(i+1)%5)
	}
	c := make(chan bool)
	<- c

}

func forkup(fork_nr, philo_nr int) {
	forks[fork_nr].Lock()
	fmt.Printf("Filozof nr %d podniósl widelec nr %d\n", philo_nr, fork_nr)

}

func forkdown(fork_nr, philo_nr int) {
	forks[fork_nr].Unlock()
	fmt.Printf("Filozof nr %d odlozyl widelec nr %d\n", philo_nr, fork_nr)
}

func philo(left, right int) {
	number := left

	for true {
		if number != 0 {
			forkup(left, number)
			time.Sleep(time.Duration(1)*time.Second)
			//czas dodany by móc zauważyć działanie
			forkup(right, number)
		} else {
			forkup(right, number)
			time.Sleep(time.Duration(1)*time.Second)
			forkup(left, number)
		}

		fmt.Printf("Filozof nr %d zaczyna jesc\n", number)
		time.Sleep(time.Duration(4)*time.Second)
		forkdown(left,number)
		forkdown(right,number)
		time.Sleep(time.Duration(4)*time.Second)
	}
}
