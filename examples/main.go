package main

import (
	"fmt"
	"github.com/dravenk/morse"
)

func main() {
	m := morse.New()

	words := "MORSE CODE!"
	code := "-- --- .-. ... .   -.-. --- -.. . -.-.--"

	c1 := m.Encode(words)
	// Output: -- --- .-. ... .   -.-. --- -.. . -.-.--
	fmt.Println(string(c1))

	m.SetSplitter("/")
	c2 := m.Encode(words)
	// Output: -- --- .-. ... . / -.-. --- -.. . -.-.-- 
	fmt.Println(c2)

	m.SetDot("0")
	m.SetDash("1")
	c3 := m.Encode(words)
	// Output: 11 111 010 000 0 / 1010 111 100 0 101011
	fmt.Println(c3)

	m.Reset()
	c4 := m.Encode(words)
	// Output: -- --- .-. ... .   -.-. --- -.. . -.-.--
	fmt.Println(c4)

	w1 := m.Decode(code)
	// Output: MORSE CODE!
	fmt.Println(w1)

	m.SetSplitter("/")
	// Input: -- --- .-. ... . / -.-. --- -.. . -.-.-- 
	w2 := m.Decode(c2)
	// Output: MORSE CODE!
	fmt.Println(w2)

	words2 := "magic MORSE CODE!"
	c5 := m.Encode(words2)
	// Output: MORSE/CODE!
	fmt.Println(c5)

	w3 := m.Decode(c5)
	// Output: MORSE/CODE!
	fmt.Println(w3)
}
