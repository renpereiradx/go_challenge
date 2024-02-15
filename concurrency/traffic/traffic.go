package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateFibo(n int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("calculando la serie de Fibonacci de %d\n", n)
	time.Sleep(2 * time.Second)
	fiboResult := Fibonacci(n)
	fmt.Printf("La serie de fibonacci de %d es: %d\n", n, fiboResult)
	<-c
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 2)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		c <- 1
		go calculateFibo(i, &wg, c)
	}
	wg.Wait()

}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
