package concurrency

// WebsiteChecker is ...
type WebsiteChecker func(string) bool

type result struct {
	url   string
	value bool
}

// CheckWebsites is ...
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	// The following will create a race condition where the same
	// bit of memory can be written to concurrently by two different invocations
	// of the goroutine
	// Use `go test -race` to detect this
	/*
		for _, url := range urls {
			// immediately invoked anonymous function called as a goroutine
			go func(currentUrl string) {
				results[currentUrl] = wc(currentUrl)
			}(url)
		}
		time.Sleep(2 * time.Second)
	*/

	// instead we define a custom struct `result`, and create a channel of `result`s
	// the channel is used to queue results from the concurrent goroutines
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(currentUrl string) {
			// add to channel whenever, can happen concurrently without creating a race
			resultChannel <- result{currentUrl, wc(currentUrl)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// read out of channel, in whatever order they happen to be added
		result := <-resultChannel
		results[result.url] = result.value
	}

	return results
}
