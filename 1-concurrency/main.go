package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func rnd(numChan chan<- int) {
	for i := 0; i < 10; i++ {
		num := rand.IntN(100)
		numChan <- num
	}
	close(numChan)
}

func squ(numChan <-chan int, resultChan chan<- int) {
	for num := range numChan {
		resultChan <- num * num
	}
	close(resultChan)
}

func main() {

	numChan := make(chan int)
	resultChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		rnd(numChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		squ(numChan, resultChan)
	}()

	go func() {
		for result := range resultChan {
			fmt.Print(result, ' ')
		}
	}()

	wg.Wait()
}
