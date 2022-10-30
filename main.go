package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gofiber/storage/memcache"
	u2 "github.com/gofiber/utils/v2"
)

func main() {
	// Initialize default config
	// if you need memcached: docker run -p 127.0.0.1:11211:11211 --name my-memcache -d memcached
	store := memcache.New()

	defer store.Close()

	_ = u2.ToLower("HI")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
