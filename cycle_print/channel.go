package cycle_print

import (
	"fmt"
	"sync"
)

func ChannelCyclePrint(times int) {
	placeHolder := struct{}{}
	channel1, channel2 := make(chan struct{}, 0), make(chan struct{}, 0)
	var w sync.WaitGroup
	w.Add(2)
	go func() {
		defer w.Done()
		for {
			_, ok := <-channel1
			if !ok {
				fmt.Println("a exit")
				return
			}
			fmt.Println("a")
			channel2 <- placeHolder
		}
	}()
	go func() {
		defer w.Done()
		for i := 0; i < times; i++ {
			_, ok := <-channel2
			if !ok {
				return
			}
			fmt.Println("b")
			if i == times-1 {
				close(channel2)
				close(channel1)
				fmt.Println("b exit")
				return
			}
			channel1 <- placeHolder
		}
	}()
	channel1 <- placeHolder
	w.Wait()
}
