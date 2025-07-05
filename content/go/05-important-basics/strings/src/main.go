package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	s := "0123456789"
	s1 := shuffle1(s)
	fmt.Println(s1)

	s2 := shuffle2(s)
	fmt.Println(s2)
}

func shuffle1(s string) string {
	runes := []rune(s)
	for i := range runes {
		j := rand.Intn(len(runes))
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func shuffle2(s string) string {
	sSlice := strings.Split(s, "")
	for i := range sSlice {
		j := rand.Intn(len(sSlice))
		sSlice[i], sSlice[j] = sSlice[j], sSlice[i]
	}
	return strings.Join(sSlice, "")
}
