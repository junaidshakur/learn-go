package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int = 0
var lock sync.Mutex

type worker struct {
	id int
}

func (w *worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d recevied %d\n", w.id, data)
		time.Sleep(time.Millisecond * 1000)
	}
}

func (w *worker) discard(c chan int) {
	for {
		data := <-c
		fmt.Println("discarding data channel", data)
		time.Sleep(time.Millisecond * 1000)
	}
}

func LearnConcurrency() {
	// fmt.Println("start")

	// for i := 0; i < 20; i++ {
	// 	// go func() {
	// 	// 	fmt.Println("processing", i)
	// 	// }()
	// 	go printLoop()
	// }

	// time.Sleep(time.Millisecond * 10)
	// fmt.Println("done")

	c := make(chan int)
	//dc := make(chan int)

	for i := 0; i < 5; i++ {
		w := &worker{id: i}
		go w.process(c)
		//go w.discard(dc)
	}

	for {
		select {
		case c <- rand.Int():
		case <-time.After(time.Millisecond * 100):
			fmt.Println("time out")
			// case dc <- rand.Int():
		default:
			fmt.Println("droped")
		}
		time.Sleep(time.Millisecond * 50)
	}
}

func printLoop() {
	lock.Lock()
	defer lock.Unlock()

	counter++
	fmt.Println("processing", counter)
}
