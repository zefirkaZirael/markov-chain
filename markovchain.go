package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	maxAllowedWords     = 10000
	defaultPrefixLength = 2
	defaultMaxWords     = 100
)

func buildMarkovChain(words []string, prefixLength int) (map[string][]string, string) {
	states := make(map[string][]string)
	if len(words) < prefixLength {
		return states, ""
	}
	prefix := strings.Join(words[:prefixLength], " ")
	for i := prefixLength; i < len(words); i++ {
		suffix := words[i]
		states[prefix] = append(states[prefix], suffix)
		prefixWords := strings.Fields(prefix)
		prefixWords = append(prefixWords[1:], suffix)
		prefix = strings.Join(prefixWords, " ")
	}
	return states, strings.Join(words[:prefixLength], " ")
}

func generateText(states map[string][]string, maxWords int, startingPrefix string, prefixLength int) string {
	var res []string

	prefix := startingPrefix
	if maxWords == 1 {
		words := strings.Fields(prefix)
		res = append(res, words[0])
		return strings.Join(res, " ")
	}
	if maxWords < prefixLength {
		words := strings.Fields(prefix)
		for i := 0; i < maxWords; i++ {
			res = append(res, words[i])
		}
		return strings.Join(res, " ")
	}

	if len(states) == 0 {
		return ""
	}

	ok := false
	var prefixes []string
	for prefix := range states {
		prefixes = append(prefixes, prefix)
		if prefix == startingPrefix {
			ok = true
		}
	}
	if !ok {
		fmt.Fprintln(os.Stderr, "Error: No such sufffix found")
		os.Exit(1)
	}

	if maxWords == 1 {
		words := strings.Fields(prefix)
		res = append(res, words[0])

		return strings.Join(res, " ")

	}
	res = append(res, prefix)

	wordCount := len(res) + prefixLength - 1
	for wordCount < maxWords {
		suffixes, exists := states[prefix]
		if !exists {
			break
		}
		suffix := suffixes[rand.Intn(len(suffixes))]

		res = append(res, suffix)

		words := strings.Fields(prefix)
		if len(words) > 0 && len(res) < maxWords {
			words = append(words[1:], suffix)
		}
		prefix = strings.Join(words, " ")
		wordCount++

	}
	return strings.Join(res, " ")
}

func printUsage() {
	fmt.Println("Markov Chain text generator.\n")
	fmt.Println("Usage:")
	fmt.Println("  markovchain [-w <N>] [-p <S>] [-l <N>]")
	fmt.Println("  markovchain --help\n")
	fmt.Println("Options:")
	fmt.Println("  --help  Show this screen.")
	fmt.Println("  -w N    Number of maximum words")
	fmt.Println("  -p S    Starting prefix")
	fmt.Println("  -l N    Prefix length")
}

func main() {
	prefix := flag.String("p", "", "Starting prefix for the Markov Chain")
	Lprefix := flag.Int("l", defaultPrefixLength, "Prefix length for the Markov Chain")
	wordCount := flag.Int("w", defaultMaxWords, "Maximum number of words to generate")
	help := flag.Bool("help", false, "Show usage information")
	flag.Parse()
	stat, _ := os.Stdin.Stat()

	//  IF NO ECHO THEN ERROR
	if *help {
		printUsage()
		os.Exit(0)
	}
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		// stdin is a terminal
		fmt.Println("Error: no input text")
		os.Exit(1)
	}

	if *Lprefix < 1 {
		fmt.Fprintln(os.Stderr, "Error: Prefix length must be at least 1")
		os.Exit(1)
	}
	if *Lprefix > 5 {
		fmt.Fprintln(os.Stderr, "Error: Prefix length cannot exceed 5")
		os.Exit(1)
	}
	if *wordCount < 0 {
		fmt.Fprintln(os.Stderr, "Error: Maximum number of words must be at least 1")
		os.Exit(1)
	}
	if *wordCount == 0 {
		fmt.Fprintln(os.Stderr, "")
		os.Exit(1)
	}
	if *wordCount > maxAllowedWords {
		fmt.Fprintln(os.Stderr, "Error: Maximum number of words cannot exceed 10,000")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var textBuilder strings.Builder

	for scanner.Scan() {
		textBuilder.WriteString(scanner.Text() + " ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	text := textBuilder.String()
	if text == "" {
		fmt.Fprintln(os.Stderr, "Error: no input text")
		os.Exit(1)
	}

	prefixLength := *Lprefix
	words := strings.Fields(text)
	if len(words) < prefixLength {
		fmt.Fprintln(os.Stderr, "Error: input text is too short for the given prefix length")
		os.Exit(1)
	}

	states, startingPrefix := buildMarkovChain(words, prefixLength)
	if *prefix != "" {
		startingPrefix = *prefix
	}

	rand.Seed(time.Now().UnixNano())

	generatedText := generateText(states, *wordCount, startingPrefix, prefixLength)
	fmt.Println(generatedText)
}
