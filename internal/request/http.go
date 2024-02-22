package request

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func Get(url string, retry int, sleep time.Duration, timeout time.Duration) (*http.Response, error) {

	client := &http.Client{Timeout: timeout}

	res, err := client.Get(url)

	if (err != nil || is5xxServerError(res.StatusCode)) && retry > 0 {

		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		sleep = time.Duration(int64(r.Intn(int(sleep)*retry) + int(1*time.Second)))

		fmt.Printf("Url: %s - Retry: %d - Sleep seconds: %v\n", url, retry, sleep)

		time.Sleep(sleep)

		retry--

		return Get(url, retry, sleep, timeout)
	}

	return res, err
}

func is5xxServerError(statusCode int) bool {
	return statusCode > 499 && statusCode < 600
}
