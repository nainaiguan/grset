package main

import (
	"fmt"
	"grset"
	"sync"
	"time"
)

func main() {
	gr := grset.NewGrSet()
	c1 := gr.Register(1)
	c2 := gr.Register(2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-c1.Done():
				fmt.Println("c1 done")
				return
			default:
				fmt.Println("1")
				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-c2.Done():
				fmt.Println("c2 done")
				return
			default:
				fmt.Println("2")
				time.Sleep(time.Second)
			}
		}
	}()

	fmt.Printf("%d routine is running\n", gr.Total())
	time.Sleep(5 * time.Second)
	gr.Shut(2)
	fmt.Printf("%d routine is running\n", gr.Total())

	wg.Wait()
}
