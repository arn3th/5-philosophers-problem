package main

import (
	"fmt"
	"time"
	"sync"
)

var forks [5]sync.Mutex
var c = make(chan struct{}, 4)

func main() {

	for i:=0; i<5; i++ {
		go philo(i,(i+1)%5)
	}

	time.Sleep(time.Duration(30)*time.Second)
	//program kończy się po 30 sekundach
}

func philo(left,right int){

	for true {
		c <- struct{}{}
		forks[left].Lock()
		fmt.Printf("Filozof nr %d podnosi lewy widelec\n",left)

		forks[right].Lock()
		fmt.Printf("Filozof nr %d podnosi prawy widelec i zaczyna jesc\n",left)

		time.Sleep(time.Duration(3)*time.Second)

		fmt.Printf("Filozof nr %d odklada prawy widelec\n",left)
		forks[right].Unlock()

		fmt.Printf("Filozof nr %d odklada lewy widelec\n",left)
		forks[left].Unlock()
		<-c
		time.Sleep(time.Duration(3)*time.Second)

	}



}