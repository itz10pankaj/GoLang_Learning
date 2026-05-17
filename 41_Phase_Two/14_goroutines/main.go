package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing Task", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	// time.Sleep(time.Second * 2)
	// The wait group is used to above above line as we do not know the how much time goroutines need to
	// exceute so we attach a waitgroup
	wg.Wait()
}
