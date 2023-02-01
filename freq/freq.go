package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Q: What is the most common word (ignoring case) in sherlock.txt
// Word frequency

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	w, err := mostCommon(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Printf(w)

}

// "Who's on first?" -> [Who s on first]
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return freqs, nil
}

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}
	return maxWord(freqs)
}

func maxWord(wordFreq map[string]int) (string, error) {
	if len(wordFreq) == 0 {
		return "", fmt.Errorf("empty map")
	}
	maxW, maxC := "", 0
	for k, v := range wordFreq {
		if v > maxC {
			maxC = v
			maxW = k
		}
	}
	return maxW, nil
}
