package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func ReadChannel(ch chan interface{}, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker was stop!")
			wg.Done()
			return
		case value, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("Value: %v \n", value)
		}
	}
}

func WriteChannel(ch chan interface{}, ctx context.Context, wg *sync.WaitGroup, values []interface{}) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			wg.Wait()
			return
		default:
			ch <- values[rand.Intn(len(values))]
		}
	}
}

func main() {
	values := []interface{}{"12", 1, 1.4, true}

	fmt.Print("Enter time in milliseconds: ")

	var seconds uint16
	if _, err := fmt.Fscan(os.Stdout, &seconds); err != nil {
		fmt.Println("Incorrect time!")
		return
	}
	if seconds == 0 {
		fmt.Println("Incorrect time!")
		return
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second * time.Duration(seconds)))
	defer cancel()

	channel := make(chan interface{})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go ReadChannel(channel, ctx, &wg)

	WriteChannel(channel, ctx, &wg, values)
}
