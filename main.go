package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

var (
	Emoji = []rune("🐜🐻🐱🐶🦅🦊🐐🦔🏝🐆🦘🦁🐵🦡🦉🐼👑🦝🐍🦃🦄🦇🐺🎸🦌🦓")
)

func ReadInput() (text string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return strings.ToLower(text)
}

func Encode(input string) string {
	output := bytes.NewBuffer(make([]byte, len(input)))
	for _, c := range input {
		if 'a' <= c && c <= 'z' {
			i := int(c - 'a')
			output.WriteRune(Emoji[(i+1)%len(Emoji)])
		} else {
			output.WriteRune(c)
		}
	}
	return output.String()
}

func main() {
	input := ReadInput()
	output := Encode(input)
	fmt.Println(output)
}
