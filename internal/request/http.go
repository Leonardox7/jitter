package request

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"
)

func Get(url string, retry int, sleep time.Duration, timeout time.Duration) (*http.Response, error) {

	client := &http.Client{Timeout: timeout}

	res, err := client.Get(url)

	if (err != nil || is5xxServerError(res.StatusCode)) && retry > 0 {

		randSleep := time.Duration(rand.IntN(int(sleep)*retry) + int(1*time.Second))

		fmt.Printf("Url: %s - Retry: %d - Sleep seconds: %v\n", url, retry, randSleep)

		time.Sleep(randSleep)

		retry--

		return Get(url, retry, sleep, timeout)
	}

	return res, err
}

func is5xxServerError(statusCode int) bool {
	return statusCode > 499 && statusCode < 600
}
