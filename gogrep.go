package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Search searches for occurences of str in a file named fileName and returns
// all matches.
func Search(fileName string, str string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matches []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), str) {
			matches = append(matches, scanner.Text())
		}
	}

	return matches
}

// SearchDir searches for occurences of str in the directory named name and
// its subdirectories and prints them.
func SearchDir(dirName string, str string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	entries, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		// ignore hidden files like .git
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		fullPath := filepath.Join(dirName, entry.Name())

		if !entry.IsDir() {
			for _, match := range Search(fullPath, str) {
				// fmt.Println(fullPath + ": " + match)
				ch <- fullPath + ": " + match
			}
		} else {
			// new goroutine to search each directory
			wg.Add(1)
			go SearchDir(fullPath, str, wg, ch)
		}
	}
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go SearchDir(".", "test", &wg, ch)

	// delegate so main can receive from ch
	go func() {
		wg.Wait()
		defer close(ch) // nothing left to receive
	}()

	numOccurrences := 0
	for occurrence := range ch {
		fmt.Println(occurrence)
		numOccurrences++
	}

	fmt.Println()
	fmt.Println(numOccurrences, "occurences")
}
