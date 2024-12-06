package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int, 6)
	array := []int{1, 2, 3, 4, 5}
	wg := sync.WaitGroup{}
	wg.Add(1) //wg.Add(len(array))
	go func() {
		for _, value := range array {
			//v := value
			ch <- &value
		}
		//close the channel close(ch)
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value)
			//wg.Done() to be outside the loop
		}
		wg.Done()
	}()

	go func() {
		var a int
		for {
			a++
		}
	}()
	wg.Wait()
}
