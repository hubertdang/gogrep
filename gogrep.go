package main

import (
	"bufio"
	"flag"
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
				ch <- fullPath + ": " + match
			}
		} else {
			// new goroutine to search each directory
			wg.Add(1)
			go SearchDir(fullPath, str, wg, ch)
		}
	}
}

// RecursiveSearch determines if the user wants to perform a recursive search.
func RecursiveSearch(rFlag string, numArgs int) bool {
	return rFlag != "" && numArgs == 0
}

// SingleSearch determines if the user wants to perform a single file search.
func SingleSearch(rFlag string, numArgs int) bool {
	return rFlag == "" && numArgs == 2
}

func main() {
	recursive := flag.String("r", "", "recursive search")
	flag.Parse()

	numArgs := flag.NArg()

	numOccurrences := 0

	if RecursiveSearch(*recursive, numArgs) {
		ch := make(chan string)
		var wg sync.WaitGroup

		wg.Add(1)
		go SearchDir(".", *recursive, &wg, ch)

		// delegate so main can receive from ch
		go func() {
			wg.Wait()
			defer close(ch) // nothing left to receive
		}()

		for occurrence := range ch {
			fmt.Println(occurrence)
			numOccurrences++
		}
	} else if SingleSearch(*recursive, numArgs) {
		for _, match := range Search(flag.Arg(0), flag.Arg(1)) {
			fmt.Println(flag.Arg(0) + ": " + match)
			numOccurrences++
		}
	} else {
		fmt.Println("Parker: [feels himself starting to disintegrate and doesn't want to")
		fmt.Println("die] Mister Stark? I don't feel so good...")
		fmt.Println()
		fmt.Println("Tony Stark: [trying to be calm] You're all right.")
		fmt.Println()
		fmt.Println("Parker: [stumbling] I don't know what's – I don't know what's happening.")
		fmt.Println("I don't – [Parker falls into Stark's arms, clutching him tight and crying]")
		fmt.Println("Save me, save me! I don't wanna go, I don't wanna go, Sir, please.")
		fmt.Println("Please, I don't wanna go. I don't wanna go... I'm sorry.")
		fmt.Println()
		fmt.Println("[Parker disintegrates into ashes in Stark's arms]")
		fmt.Println()
		fmt.Println("Nebula: [to Tony Stark, seeing Thanos' victory] He did it.")
		fmt.Println()
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(numOccurrences, "occurences")
}
