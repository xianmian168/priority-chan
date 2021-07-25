package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
)

func handle(low, high <-chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case highJob := <-high:
			log.Println("highJob: ", highJob)
		case lowJob := <-low:
		p:
			for {
				select {
				case highJob := <-high:
					log.Println("force highJob: ", highJob)
				default:
					break p
				}
			}
			log.Println("lowJob: ", lowJob)
		}
	}
}
func main() {
	high := make(chan int, 10)
	low := make(chan int, 10)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	go func() {
		for i := 0; i < 10; i++ {
			high <- i
		}
	}()

	go func() {
		for i := 0; i < 2; i++ {
			low <- i
		}
	}()

	time.Sleep(time.Second * 1)
	go handle(low, high, ctx)
	time.Sleep(time.Second * 10)
	cancel()
}
