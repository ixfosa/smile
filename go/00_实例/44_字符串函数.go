package main

import (
	"fmt"
	s "strings"
)

func main() {
	var p = fmt.Println

	p("Contains: ", s.Contains("test", "es")) //Contains: true

	p("Count: ", s.Count("test", "t")) // Count: 2

	p("HasPrefix: ", s.HasPrefix("test", "te")) // HasPrefix:  true

	p("HasSuffix: ", s.HasSuffix("test", "st")) // HasSuffix:  true

	p("Index:     ", s.Index("test", "e")) // Index: 1

	p("Join:      ", s.Join([]string{"a", "b"}, "-")) //Join: a-b

	p("Repeat:    ", s.Repeat("a", 5)) // Repeat: aaaaa

	p("Replace:   ", s.Replace("foo", "o", "0", -1)) // Replace: f00

	p("Replace:   ", s.Replace("foo", "o", "0", 1)) // Replace: f0o

	p("Split:     ", s.Split("a-b-c-d-e", "-")) // Split: [a b c d e]

	p("ToLower:   ", s.ToLower("TEST")) // toLower: test

	p("ToUpper:   ", s.ToUpper("test")) //ToUpper: TEST

	p("Len: ", len("hello")) // Len:  5

	p("Char:", "hello"[1]) // Char: 101
}
