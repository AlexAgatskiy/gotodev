package main

import (
	"fmt"
	"time"
)

// Вопрос на ковариативность / инвариативность типов (определения из ООП,
// полезно знать) - есть ли она в Го? Или, на примере, что будет напечатано?
// (тип time реализует интерфейс fmt.Stringer).

func f1(sl []fmt.Stringer) {
	for _, s := range sl {
		fmt.Println(s)
	}
}

func main() {
	times := []time.Time{time.Now(), time.Now()}
	f1(times)
}
