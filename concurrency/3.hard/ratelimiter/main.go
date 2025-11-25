package main

import (
	"fmt"
	"time"
)

// Есть функция, которая выполняет http-запрос. Необходимо реализовать логику
// rate limiter с возможностью ограничения кол-ва запросов, производимых этой
// функцией до n-го кол-ва в секунду.

func makeRequest(url, data string) error {
	// Здесь могла быть реальная работа HTTP-клиента.
	fmt.Println("Request sent to:", url, "with data:", data, "at",
		time.Now().Format("15:04:05.000"))
	return nil
}

func makeRequestWithLimit() error {

}

func main() {
	// ?? rl := NewRateLimiter(3) ??

	for i := 0; i < 10; i++ {
		go makeRequestWithLimit()
	}

	time.Sleep(3 * time.Second)
}
