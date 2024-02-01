package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 0.1}
	set := make(map[int64][]float64)
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for _, v := range temp {
		wg.Add(1)
		go func(value float64) {
			defer wg.Done()

			var dozen float64

			if value < 0 {
				dozen = math.Ceil(value / 10)
			} else {
				dozen = math.Ceil(value/10) - 1
			}

			key, _ := strconv.ParseInt(fmt.Sprintf("%d0", int64(dozen)), 0, 64)

			mutex.Lock()
			set[key] = append(set[key], value)
			mutex.Unlock()
		}(v)
	}

	wg.Wait()

	for index, value := range set {
		fmt.Printf("%d: %.1f\n", index, value)
	}
}
