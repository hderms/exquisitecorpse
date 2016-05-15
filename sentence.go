package main

import (
	"fmt"
	"strings"
)

type Sentence struct {
	words     []string
	max_words int
}

func NewSentence() *Sentence {

	return &Sentence{max_words: 5}
}

func (s *Sentence) AddWord(word string) {
	s.words = append(s.words, word)
}

func (s *Sentence) HandleText(text string) bool {

	fmt.Printf("Received text: %s\n", text)

	if len(s.words) < s.max_words {
		text = strings.TrimSpace(text)
		s.AddWord(text)
		fmt.Printf("New senetence is: %s\n", strings.Join(s.words, " "))
		return true
	} else {
		return false
	}
}
