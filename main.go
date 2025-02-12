package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var (
	Letters = []rune("abcdefghijklmnopqrstuvwxyz")
	Emoji   = []rune("🐜🐻🐱🐶🦅🦊🐐🦔🏝🐆🦘🦁🐵🦡🦉🐼👑🦝🐍🦃🦄🦇🐺🎸🦌🦓")
)

func ReadInput() (text string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return strings.ToLower(text)
}

func main() {
	input := ReadInput()
	output := bytes.NewBuffer(make([]byte, len(input)))
	for _, c := range input {
		var found bool
		for i, l := range Letters {
			if l == c {
				i++
				if i > len(Letters) {
					i = 0
				}
				output.WriteRune(Emoji[i])
				found = true
				break
			}
		}
		if !found {
			output.WriteRune(c)
		}
	}
	fmt.Println(output.String())
}
