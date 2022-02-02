package main

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
)

func main() {
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	u := []rune(uppercase)
	l := len(u)
	res := make(chan string, l*2)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < l+2; i += 2 {
		wg.Add(1)
		go func(ticker int, result chan string) {
			defer wg.Done()
			mu.Lock()
			var buffer bytes.Buffer
			d1 := strconv.Itoa(ticker + 1)
			d2 := strconv.Itoa(ticker + 2)
			buffer.WriteString(d1)
			buffer.WriteString(d2)
			result <- fmt.Sprint(buffer.String())
			mu.Unlock()
		}(i, res)
		wg.Wait()
		wg.Add(1)
		go func(ticker int, result chan string, sequence []rune) {
			defer wg.Done()
			mu.Lock()
			if ticker < l {
				result <- string(sequence[ticker])
				result <- string(sequence[ticker+1])
			}
			mu.Unlock()
		}(i, res, u)
		wg.Wait()
	}
	close(res)
	for {
		val, ok := <-res
		if ok == false {
			fmt.Printf("%v", val)
			break // exit break loop
		} else {
			fmt.Printf("%v", val)
		}
	}
}
