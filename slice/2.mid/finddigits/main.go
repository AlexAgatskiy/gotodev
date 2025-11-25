package main

import (
	"fmt"
	"os"
	"regexp"
)

// Подумать, что не так с функцией findDigits с точки зрения работы с памятью.
// То, что функция игнорирует ошибку или не проверяет файл на существование, нас
// не интересует. Подсказка в 25-ой строке.

var digitRegexp = regexp.MustCompile("[0-9]+")

func findDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)
	return digitRegexp.Find(b)
}

func main() {
	result := findDigits("some_file_path")

	fmt.Println(result)

	// Дальнейшая работа с result...
}
