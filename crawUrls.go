package main

import (
	"fmt"
	"strconv"
	"time"
)

func initDataUrls(length int) []string {
	sliceUrls := make([]string, 0)
	for i := 0; i < length; i++ {
		sliceUrls = append(sliceUrls, "data url "+strconv.Itoa(i))
	}
	return sliceUrls
}

func crawDataUrl(url string, maxConcurrentChannel chan string, keepMainAliveChannel chan string) {
	fmt.Printf("craw data url...: %s\n", url)
	time.Sleep(5000 * time.Millisecond)
	receive := <-maxConcurrentChannel
	keepMainAliveChannel <- receive
}

func main() {
	length := 1000
	urls := initDataUrls(length)

	maxConcurrentChannel := make(chan string, 5)
	keepMainAliveChannel := make(chan string, length)

	for _, v := range urls {
		//khi số lượng goroutine đang chạy là 5,hàm main sẽ bị lock tại đây
		maxConcurrentChannel <- v
		go crawDataUrl(v, maxConcurrentChannel, keepMainAliveChannel)
	}

	for i := 0; i < length; i++ {
		//khi vòng for trên chạy xong,5 goroutine cuối vẫn đang chạy
		//hàm main sẽ bị block nếu keepMainAliveChannel bị rỗng,
		//đảm bảo 5 goroutine cuối cùng chạy xong thì hàm main mới kết thúc
		<-keepMainAliveChannel
	}
}
