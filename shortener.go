package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

var words []string
var directory map[string]string
var directoryRev map[string]string

func readFile() {

	f, _ := os.Open("./bip39.txt")
	words = make([]string, 2048)
	scanner := bufio.NewScanner(f)
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		words[i] = line
		i++
	}
}

func getRandomWords(count int) string {

	rand.Seed(time.Now().Unix())
	str := ""

	for i := 0; i < count; i++ {
		index := rand.Intn(len(words))
		str += words[index]
	}

	return str

}
func isValidURL(tocheck string) bool {
	_, err := url.ParseRequestURI(tocheck)
	return err == nil
}

func getMappedURL(randomwords string) string {
	return directory[strings.ToLower(randomwords)]
}

func mapURL(url string) string {

	// if this URL has already been saved before
	if directoryRev[url] != "" {
		return directoryRev[url]
	}

	if isValidURL(url) {

		randomwords := getRandomWords(3)

		for directory[randomwords] != "" {
			randomwords = getRandomWords(3)
		}

		directory[randomwords] = url
		directoryRev[url] = randomwords
		return randomwords
	}
	fmt.Println("URL is not valid")
	return ""
}
