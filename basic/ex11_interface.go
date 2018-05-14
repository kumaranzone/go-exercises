package main

import "fmt"

// VowelsFinder has FindVowels method
type VowelsFinder interface {
	FindVowels() []rune
}

// MyString - string type
type MyString string

// FindVowels implements MyString
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Kumaran")
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c", v.FindVowels())
}
