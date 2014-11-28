package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type result struct {
	url, body string
	depth     int
	urls      []string
	err       error
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// This implementation doesn't do either:
	results := make(chan result)
	quit := make(chan int)
	fetched := make(map[string]bool)
	fetching := 0

	defer close(results)
	defer close(quit)

	// Fetch URLs in parallel.
	go fetch(url, depth, fetcher, results, quit)
	fetched[url] = true
	fetching += 1

	for {
		select {
		case res := <-results:
			if res.err != nil {
				fmt.Println(res.err)
				break
			}
			fmt.Printf("found: %s %q\n", res.url, res.body)
			for _, u := range res.urls {
				// Don't fetch the same URL twice.
				if !fetched[u] {
					// Fetch URLs in parallel.
					go fetch(u, res.depth-1, fetcher, results, quit)
					fetched[u] = true
					fetching += 1
				}
			}
		case <-quit:
			if fetching -= 1; fetching == 0 {
				return
			}
		}
	}
}

func fetch(url string, depth int, fetcher Fetcher, results chan result, quit chan int) {
	if depth <= 0 {
		quit <- 0
		return
	}
	body, urls, err := fetcher.Fetch(url)
	results <- result{url, body, depth, urls, err}
	quit <- 0
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
