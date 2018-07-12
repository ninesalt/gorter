package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/go-redis/redis"
)

var words []string
var redisClient *redis.Client

func connectToRedis() {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

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

// get URL given words
func getMappedURL(randomwords string) string {

	URL, newErr := redisClient.Get(randomwords).Result()

	if newErr == nil {
		return URL
	}

	return ""

}

func mapURL(url string) string {

	// if this URL has already been saved before
	words, newErr := redisClient.Get(url).Result()

	if newErr == nil {
		return words
	}

	if isValidURL(url) {

		randomwords := getRandomWords(3)
		initial, _ := redisClient.Get(randomwords).Result()

		for initial != "" {
			randomwords = getRandomWords(3)
			initial, _ = redisClient.Get(randomwords).Result()
		}

		redisClient.Set(randomwords, url, 0)
		redisClient.Set(url, randomwords, 0)
		return randomwords
	}
	fmt.Println("URL is not valid")
	return ""
}
