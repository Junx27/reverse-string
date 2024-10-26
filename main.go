package main

import "fmt"

func ReverseString(s string) string {
	result := []rune(s)
	reversed := make([]rune, len(result))

	for i := 0; i < len(result); i++ {
		reversed[i] = result[len(result)-1-i]
	}
	return string(reversed)
}

func main() {
	word := "tri saptono"
	fmt.Println(ReverseString(word))
}
