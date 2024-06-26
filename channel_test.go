/*
[Channel]
Channel berfungsi untuk komunikasi antar goroutine
Misal butuh menerima data dari proses goroutine
*/
package golang_goroutine

import (
	"fmt"
	"strconv"
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

// ----------------------------------------------------------------

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// Tanpa menggunakan goroutine
	// channel <- "Fadli"
	// channel <- "Darusalam"
	// channel <- "Sragen"

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// Dengan Goroutine
	go func() {
		channel <- "Fadli"
		channel <- "Darusalam"
		channel <- "Sragen"
	}()
	
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai!")
}

// ----------------------------------------------------------------

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

// ----------------------------------------------------------------

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

// ----------------------------------------------------------------

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <- channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}

		if counter == 2 {
			break
		}
	}
}