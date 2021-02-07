package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	receiveSignal()
}

func receiveSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	fmt.Println("Waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}

func cancelContext() {
	fmt.Println("start main")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("Context sub() is running")
		time.Sleep(time.Second)
		fmt.Println("Context sub() is finished")
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("end main")
}

func printPrimeNumber() {
	pn := primeNumber()
	for n := range pn {
		fmt.Println(n)
	}
}

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j <= l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func channel() {
	fmt.Println("start main")
	done := make(chan bool)
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
		done <- true
	}()
	<-done
	fmt.Println("end main")
}

func inline() {
	fmt.Println("start main")
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("end main")
}

func callSub() {
	fmt.Println("start main")
	go sub()
	time.Sleep(2 * time.Second)
	fmt.Println("end main")
}

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}
