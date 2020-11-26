package main

import "fmt"
import "errors"
import "net/http"

var errRequsetFailed = errors.New("Request failed")

func main() {
	// map need to initialize
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.naver.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}

}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequsetFailed
	}
	return nil

}
