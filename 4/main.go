package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func Worker(ctx context.Context, name int, ch chan interface{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker #%d shutdown\n", name+1)
			wg.Done()
			return
		//забираем данные из канала
		case value, ok := <-ch:
			if !ok {
				break
			}
			fmt.Printf("Worker #%d, Value: %v\n", name+1, value)
		}
	}
}

func main() {
	//данные для канала
	array := []interface{}{"2312", 34, 45, 0.3, false}
	channel := make(chan interface{})

	fmt.Print("Enter a number workers: ")

	var number uint
	if _, err := fmt.Fscan(os.Stdin, &number); err != nil {
		fmt.Println("Value is not a number!")
		return
	}
	wg := sync.WaitGroup{}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	for i := 0; i < int(number); i++ {
		wg.Add(1)
		go Worker(ctx, i, channel, &wg)
	}

	for {
		select {
		case <-ctx.Done():
			close(channel)
			wg.Wait()
			return
			//записываем данные в канал
		default:
			channel <- array[rand.Intn(len(array))]
		}
	}
}
