package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	easyMap := make(map[string]int)
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for i := 1; i < 100; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			key := strconv.Itoa(index)

			mutex.Lock()
			easyMap[key] = index
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	for index, value := range easyMap {
		fmt.Printf("Index = %s, Value = %d\n", index, value)
	}
}
