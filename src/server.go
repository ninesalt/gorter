package main

import (
	"log"
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	requested := r.URL.Path

	// serve index.html
	if requested == "/" {
		http.ServeFile(w, r, "./web/index.html")

	} else if strings.HasSuffix(requested, ".css") { //serve requested CSS files
		http.ServeFile(w, r, "./web/"+requested)

	} else {
		// if the path is /randomwords, redirect to mapped URL
		randomwords := r.URL.Path[1:]
		url := getMappedURL(randomwords)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

// get random words given URL in GET request
func getWords(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		url := r.URL.Query()["url"][0]
		mapping := mapURL(url)

		// if URL cannot be mapped, add an error status code
		if mapping == "" {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Write([]byte(mapping))
	}
}

func main() {

	readFile()
	connectToRedis()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/map", getWords)
	log.Println("listening on port 5000")
	http.ListenAndServe(":5000", nil)

}
