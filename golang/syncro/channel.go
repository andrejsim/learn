package main

import (
	"log"
	"sync"
)

type on struct {
	res chan int
}

type incrementOp struct {
	on
}

type getValueOp struct {
	on
}

func newIncrementOp() incrementOp {
	return incrementOp{
		on: on {
			res: make(chan int),
		},
	}
}

func newGetValueOp() getValueOp {
	return getValueOp{
		on: on {
			res: make(chan int),
		},
	}
}

func increment(ops chan<- incrementOp, wg *sync.WaitGroup) {
	defer wg.Done()

at: = newIncrementOp ()
	ops <- op
	<-op.res
}

func main() {
	incrementOps := make(chan incrementOp)
	getValueOps := make(chan getValueOp)

	go func() {
		counter := 0
		for {
			select {
			case op := <-incrementOps:
				counter++
				op.res <- counter
			case op := <-getValueOps:
				op.res <- counter
			}
		}
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(2)

		go increment(incrementOps, &wg)
		go increment(incrementOps, &wg)
	}

	wg.Wait()

	getValueOp := newGetValueOp()
	getValueOps <- getValueOp

	log.Printf("Counter: %d", <-getValueOp.res)
}