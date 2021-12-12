package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions in Golang")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Routine 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Routine 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Routine 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Routine 4")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}
