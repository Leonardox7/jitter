package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Leonardox7/jitter/internal/request"
)

func main() {
	res, err := request.Get("https://www.google.com/robots.txt", 3, 2*time.Second, 2*time.Microsecond)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	fmt.Printf("Result: %s\n", body)
}
