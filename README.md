# Roulette

Roulette is a module designed to provide random user-agents and referrer headers.

## Usage

``` go
package main

import (
	"fmt"
	"github.com/zer-far/roulette"
)

func main() {
	// Get a random user agent
	userAgent := roulette.GetUserAgent()
	fmt.Println("Random User Agent:", userAgent)

	// Get a random referrer
	referrer := roulette.GetReferrer()
	fmt.Println("Random Referrer:", referrer)
}
```

``` shell
go run main.go
```
