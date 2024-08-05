package main

import (
	"fmt"
	"time"
)

func produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consume(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int, 5) // Buffered channel with a size of 5

	go produce(ch)
	go consume(ch)
	
	time.Sleep(1 * time.Second)
}

/*
4.2. Perbedaan Perilaku Antara Channels Buffered dan Unbuffered:
Buffered Channels: Goroutine pengirim memblokir hingga goroutine penerima menerima nilainya. 
Artinya pengiriman dan penerimaan terjadi secara bersamaan.
Unbuffered Channels: Goroutine pengirim tidak akan diblokir sampai buffernya penuh. 
Hal ini memungkinkan goroutine pengirim untuk melanjutkan eksekusi hingga kapasitas buffer tercapai. 
Goroutine penerima dapat menerima nilai sesuai kecepatannya sendiri. 
Hal ini dapat meningkatkan kinerja dengan memisahkan operasi pengiriman dan penerimaan, namun memerlukan pengelolaan ukuran buffer dengan benar untuk menghindari pengisian berlebih.
*/