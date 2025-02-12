package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

var (
	Emoji = []rune("🐜🐻🐱🐶🦅🦊🐐🦔🏝🐆🦘🦁🐵🦡🦉🐼👑🦝🐍🦃🦄🦇🐺🎸🦌🦓")
	Size  = len(Emoji)
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
			output.WriteRune(Emoji[(i+1)%Size])
		} else {
			output.WriteRune(c)
		}
	}
	return output.String()
}

func Decode(input string) string {
	output := bytes.NewBuffer(make([]byte, len(input)))
	for _, c := range input {
		if i := slices.Index(Emoji, c); i > -1 {
			i = (i - 1 + Size) % Size
			output.WriteRune(rune(i + 'a'))
		} else {
			output.WriteRune(c)
		}
	}
	return output.String()
}

func main() {
	input := ReadInput()
	var output string
	if len(os.Args) > 1 && os.Args[1] == "-d" {
		output = Decode(input)
	} else {
		output = Encode(input)
	}
	fmt.Println(output)
}
