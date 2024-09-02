package osCommands

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func Greeter() {
	p1 := os.Args[1]
	p2 := os.Args[2]

	fmt.Printf("\"Hello, %s\"...says %s\n", p1, p2)
	fmt.Printf("len: %d\n", utf8.RuneCountInString(p1))
}

func Banger() {
	word := os.Args[1]

	l := utf8.RuneCountInString(word)
	b := strings.Repeat("!", l)
	res := b + strings.ToUpper(word) + b

	fmt.Println(res)
}
