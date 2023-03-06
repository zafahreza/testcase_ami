package main

import (
	"fmt"
	"sync"
)

func check(i int, numberChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	if i%2 != 0 {
		numberChan <- i
	}
}

func print(numberChan chan int) {
	for {
		fmt.Println(<-numberChan)
	}
}

func main() {

	numberChan := make(chan int)
	wg := sync.WaitGroup{}

	go print(numberChan)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go check(i, numberChan, &wg)
	}

	wg.Wait()
	close(numberChan)

}
