package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
	notes:
	- ketika ada kirim data ke channel tapi tidak ada yang menerima, maka error: "panic: send on closed channel"
	- ketika ada menerima data dari channel tapi tidak ada yang mengirim, maka error time out atau deadlock
*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Muhammad Nafi Furqon Diani"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println("data", data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Nafi Furqon Diani"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("data", data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	/*
		error karena sudah ditandai hanya menerima channel untuk mengirim data
		invalid operation: cannot receive from send-only channel channel (variable of type chan<- string)
	*/
	// data := <-channel
	channel <- "Muhammad Nafi Furqon Diani"
}

func OnlyOut(channel <-chan string) {
	/*
		error karena sudah ditandai hanya menerima channel untuk menerima data
		invalid operation: cannot send to receive-only channel channel (variable of type <-chan string)
	*/
	// channel <- "Muhammad Nafi Furqon Diani"
	data := <-channel
	fmt.Println("data", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Muhammad"
		channel <- "Nafi"
		channel <- "Furqon"
	}()

	go func() {
		fmt.Println(cap(channel))
		fmt.Println(len(channel))
		fmt.Println(<-channel)
		fmt.Println("==================")
		fmt.Println(cap(channel))
		fmt.Println(len(channel))
		fmt.Println(<-channel)
		fmt.Println("==================")
		fmt.Println(cap(channel))
		fmt.Println(len(channel))
		fmt.Println(<-channel)
		fmt.Println("==================")
		fmt.Println(cap(channel))
		fmt.Println(len(channel))
		fmt.Println("==================")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}
