package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}
func manyWords(words ...string) {
	fmt.Println(words)
}
func lenAndUpper(char string) (length int, uppercase string) {
	defer fmt.Println("Finish The Process")
	length = len(char)
	uppercase = strings.ToUpper(char)
	return
}
func main() {
	manyWords("Aldo", "Ryanda ", "John", "Doe")
	totalLengthName, charName := lenAndUpper("Aldo")
	fmt.Println(totalLengthName, charName)
}
