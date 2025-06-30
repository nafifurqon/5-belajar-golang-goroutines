package belajar_golang_goroutines

import (
	"fmt"
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
