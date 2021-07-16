package main

import (
	"errors"
	"log"
	"sync"
)

type SafeNumbers struct {
	sync.RWMutex
	numbers map[int]int
}

func (sn *SafeNumbers) Add(num int) {
	sn.Lock()
	defer sn.Unlock()
	sn.numbers[num] = num
}

func (sn *SafeNumbers) Get(num int) (int, error) {
	sn.RLock()
	defer sn.RUnlock()
	if number, ok := sn.numbers[num]; ok {
		return number, nil
	}
	return 0, errors.New("Number does not exists")
}

func main() {
	generateNumbersMap()
}

func generateNumbersMap() {
	wg := sync.WaitGroup{}
	safeNumbers := &SafeNumbers{
		numbers: map[int]int{},
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeNumbers.Add(i)
		}(i)
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			number, err := safeNumbers.Get(i)
			if err != nil {
				log.Print(err)
			} else {
				log.Print(number)
			}
		}(i)
	}

	wg.Wait()
}
