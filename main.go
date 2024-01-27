package roulette

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	userAgents    []string
	userAgentsLen int
	rng           *rand.Rand
	mu            sync.Mutex
)

var domainName = []string{
	"google",
	"duckduckgo",
	"yahoo",
	"bing",
	"baidu",
	"yandex",
	"reddit",
	"pinterest",
	"quora",
	"twitter",
	"facebook",
}

var tld = []string{
	"com",
	"net",
	"co",
	"org",
}

var (
	domainLen = len(domainName)
	tldLen    = len(tld)
)

func init() {
	randSource := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(randSource)

	if err := fetchUserAgents(); err != nil {
		fmt.Println("Error fetching user agents:", err)
	}
}

// Fetch user agents
func fetchUserAgents() error {
	resp, err := http.Get("https://cdn.jsdelivr.net/gh/microlinkhq/top-user-agents@master/src/index.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	mu.Lock()
	defer mu.Unlock()

	if err := json.Unmarshal(body, &userAgents); err != nil {
		return err
	}

	userAgentsLen = len(userAgents)
	return nil
}

// Return random user agent
func getRandomUserAgent() string {
	mu.Lock()
	defer mu.Unlock()

	if userAgentsLen == 0 {
		return ""
	}

	return userAgents[rng.Intn(userAgentsLen)]
}

// Return random referrer
func GenerateRandomReferer() string {
	mu.Lock()
	defer mu.Unlock()

	selectedDomain := domainName[rand.Intn(domainLen)]
	selectedTLD := tld[rand.Intn(tldLen)]

	return fmt.Sprintf("https://www.%s.%s", selectedDomain, selectedTLD)
}
