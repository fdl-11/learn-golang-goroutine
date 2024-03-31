/*
Channel berfungsi untuk komunikasi antar goroutine
Misal butuh menerima data dari proses goroutine
*/
package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// kirim data / memasukkan data ke channel = channel <- data
	// channel <- "Fadli"
	
	// menerima data / mengambil data dari channel = data <- channel
	// data := <- channel
	
	// var data string
	// data = <- channel
	
	// data dari channel langsung dikirim ke parameter. Parameter mengambil data dari channel
	// fmt.Println(<- channel)
	
	// Jangan lupa channel di close
	// close(channel)			// taruh akhir
	// defer close(channel)		// taruh di awal bisa

	channel := make(chan string)
	defer close(channel)

	// Goroutine with anonymous function || jarang digunakan
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Fadli Darusalam"
		fmt.Println("Selesai mengirim data ke channel")
	}()
	
	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// ----------------------------------------------------------------

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Fadli Darusalam"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// ----------------------------------------------------------------

// Channel yg hanya untuk mengirim data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Fadli Darusalam"
}

// Channel yg hanya untuk menerima data
func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}