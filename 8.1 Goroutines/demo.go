package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	go spinner(100 * time.Millisecond)
	const n = 40
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	end := time.Since(t1)
	fmt.Println(end)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}
