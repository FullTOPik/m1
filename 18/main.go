package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

func Increment(name string, ctx context.Context, c *Counter) {
	for {
		select {
		case <-ctx.Done():
			c.wg.Done()
			return
		default:
			c.mutex.Lock()
			c.Value++
			fmt.Printf("%s, Value = %d\n", name, c.Value)
			c.mutex.Unlock()
		}
	}
}

// Конкуретный доступ реализован с помощью мьютекса

type Counter struct {
	wg    sync.WaitGroup
	mutex sync.Mutex
	Value int64
}

func main() {
	counter := Counter{
		wg:    sync.WaitGroup{},
		mutex: sync.Mutex{},
		Value: 0,
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	for i := 0; i < 5; i++ {
		counter.wg.Add(1)
		name := fmt.Sprintf("worker%d", i)
		go Increment(name, ctx, &counter)
	}

	counter.wg.Wait()

	fmt.Printf("Result value = %d", counter.Value)
}
