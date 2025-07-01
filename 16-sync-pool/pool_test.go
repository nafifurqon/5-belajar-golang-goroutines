package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPoolWithoutDelay(t *testing.T) {
	pool := sync.Pool{}

	pool.Put("Nafi")
	pool.Put("Furqon")
	pool.Put("Diani")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("data:", data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
}

func TestPoolWithDelay(t *testing.T) {
	pool := sync.Pool{}

	pool.Put("Nafi")
	pool.Put("Furqon")
	pool.Put("Diani")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("data:", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}

func TestPoolWithDelayAndNew(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Nafi")
	pool.Put("Furqon")
	pool.Put("Diani")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("data:", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}

func TestPoolWithWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}
	pool := sync.Pool{}

	pool.Put("Nafi")
	pool.Put("Furqon")
	pool.Put("Diani")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println("data:", data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Selesai")
}
