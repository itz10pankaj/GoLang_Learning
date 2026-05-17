package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wq *sync.WaitGroup) {
	defer func() {
		wq.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.views++
}

func main() {
	var wq sync.WaitGroup
	Mypost := post{
		views: 0,
	}

	for i := 0; i < 100; i++ {
		wq.Add(1)
		go Mypost.inc(&wq)
	}
	wq.Wait()
	fmt.Println(Mypost.views)
}
