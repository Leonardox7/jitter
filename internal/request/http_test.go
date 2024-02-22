package request_test

import (
	"testing"
	"time"

	"github.com/Leonardox7/jitter/internal/request"
)

func TestGet_should_return_error_timeout(t *testing.T) {
	_, err := request.Get("https://www.google.com/robots.txt", 2, 2*time.Second, 1*time.Microsecond)
	if err == nil {
		t.Errorf("expetected error but receive nil")
	}
}

func TestGet_should_return_error_nil(t *testing.T) {
	_, err := request.Get("https://www.google.com/robots.txt", 3, 1*time.Second, 2*time.Second)
	if err != nil {
		t.Errorf("got %v, wanted %v", err, nil)
	}
}
