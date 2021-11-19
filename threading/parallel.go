package threading

import (
	"log"
	"sync"
)

type Parallel struct {
	parallelDegree int
	await          *sync.WaitGroup
	watchDog       chan struct{}
}

func (lib *Parallel) New(parallelDegree int) {
	lib.await = &sync.WaitGroup{}
	lib.parallelDegree = parallelDegree
	lib.watchDog = make(chan struct{}, parallelDegree)
}

func (lib Parallel) Run(action func(interface{}), parameters interface{}) {

	lib.await.Add(1)
	lib.watchDog <- struct{}{}

	go func(function func(interface{}), params interface{}) {

		defer func() {
			if r := recover(); r != nil {
				log.Println("FATAL: ", r)
			}
		}()

		function(params)

		<-lib.watchDog
		lib.await.Done()
	}(action, parameters)
}

func (lib Parallel) Wait() {
	lib.await.Wait()
}
