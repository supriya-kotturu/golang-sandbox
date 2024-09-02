package basics

import (
	"fmt"
	"sync"
)

func even(wg *sync.WaitGroup, evenCh chan<- int, oddCh <-chan int) {
	defer func() {
		wg.Done()
		close(evenCh)
	}()

	for i := 0; i < 20; i += 2 {
		signal := <-oddCh
		fmt.Println(i)
		evenCh <- signal
	}
}

func odd(wg *sync.WaitGroup, evenCh <-chan int, oddCh chan<- int) {
	defer func() {
		wg.Done()
		close(oddCh)
	}()

	for i := 1; i < 20; i += 2 {
		signal := <-evenCh
		fmt.Println(i)
		oddCh <- signal
	}
}

func EvenOdd() {
	fmt.Println("EVEN_ODD")

	evenCh := make(chan int, 1)
	oddCh := make(chan int, 1)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go odd(&wg, evenCh, oddCh)
	go even(&wg, evenCh, oddCh)

	oddCh <- 1

	wg.Wait()
}
