package main

import "fmt"

// Что будет выведено? Почему? Подумать, после запустить программу. Исправить.

func main() {
	var maxi int

	for i := 1000; i > 0; i-- {
		go func() {
			if i%2 == 0 && i > maxi {
				maxi = i
			}
		}()
	}

	fmt.Printf("Maximum is %d", maxi)
}
