// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing(ch chan int, done chan bool) {
	//TODO: increment i 1000000 times
	for j:=0; j<1_000_000; j++{
		ch <- 1
	}
	done <- true // signaling that the go ruitine is done
}

func decrementing(ch chan int, done chan bool) {
	//TODO: decrement i 1000000 times

	for j:=0; j<1_000_000+1; j++{
		ch <- -1
	}
	done <- true // signaling that the go ruitine is done

}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)

	// TODO: Spawn both functions as goroutines
	ch := make(chan int)
	done := make(chan bool)

	go incrementing(ch, done)
	go decrementing(ch, done)

	go func() {
		for{
			i += <- ch
		}
	}()

	<- done
	<- done

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	Println("The magic number is:", i)
}
