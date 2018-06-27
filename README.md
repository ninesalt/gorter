<div align = "center">

<img src = "https://image.ibb.co/fZf9h8/ia.png" with="250" height="250"/>

</div>

## What is Gorter?

Gorter is a fast and simple URL shortener built using Go. It simpy maps a URL to 3 random [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) words that are easy to read and remember. 

To run the server just run both `.go` files:

`go run .\shortener.go .\server.go`

This will run a fresh server on port 5000 and will also serve the frontend interface. 

## API
To map a URL, send a get request to `/map` with a valid URL in the `url` header. The server will respond with the random words that map to this URL. So use them navigate to `/yourphrasehere` and you should be redirected to your URL. 