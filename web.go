package helper

import "net/http"

// Takes input string url and returns status_code and error
func GetStatus(url string) (int, error) {
	re, re_err := http.Get(url)
	return re.StatusCode, re_err
}
