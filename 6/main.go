package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	//получением сообщения
	go func() {
		wg.Add(1)
		fmt.Println("First was start!")
		for {
			select {
			case <-ch:
				fmt.Println("First was stop!")
				wg.Done()
				return
			default:
				fmt.Println("First goroutine working!")
				time.Sleep(time.Second / 2)
			}
		}
	}()
	time.Sleep(1 * time.Second)

	ch <- 1

	//закрытие горутины
	wg.Add(1)
	go func() {
		fmt.Println("\nSecond was start!")
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					fmt.Println("Second was stop!")
					wg.Done()
					return
				}
				fmt.Println("Second goroutine working!")
				time.Sleep(time.Second / 2)
			}
		}
	}()

	ch <- 1
	ch <- 2
	close(ch)

	//таймаут
	timer := time.After(time.Second)

	wg.Add(1)
	go func() {
		fmt.Println("\nThird was start!")
		for {
			select {
			case <-timer:
				fmt.Println("Third was stop!")
				wg.Done()
				return
			default:
				fmt.Println("Third goroutine working!")
				time.Sleep(time.Second / 2)
			}
		}
	}()

	//Передача контекста
	ctx, stop := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(2)*time.Second))
	defer stop()

	wg.Add(1)
	go func() {
		fmt.Println("\nFourth was start!")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Fourth was stop!")
				wg.Done()
				return
			default:
				fmt.Println("Fourth goroutine working!")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}
