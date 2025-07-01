package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	/*
		kalau begini saja akan error time out
		karena di function WaitCondition() ada cond.Wait() atau meminta menunggu
		tapi tidak ada cond.Signal() atau memberi tahu untuk lanjut/berhenti menunggu
	*/
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	/*
		pakai cond.Signal() untuk memberi tahu goroutine yang sedang muenunggu untuk lanjut
	*/
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()

	/*
		- pakai cond.Broadcast() untuk memberi tahu semua goroutine yang sedang muenunggu untuk lanjut
		- pakai goroutine supaya sleep selama 1 detiknya berjalan secara paralel dengan goroutine lain
		yang sedang menjalankan WaitCondition()
	*/

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}
