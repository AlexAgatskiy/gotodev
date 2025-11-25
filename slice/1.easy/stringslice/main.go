package main

import "fmt"

// Что будет выведено? Почему? Подумать, после запустить программу.

func main() {
	sl := make([]string, 5, 7)
	f1(sl)
	fmt.Println(sl)
}

func f1(sl []string) {
	sl = append(sl, "Hello World!")
}
