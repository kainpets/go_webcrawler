package main

import (
	"fmt"
	"github.com/kainpets/go_webcrawler/normalize_url"
)

func main() {
	url, err := normalize_url.NormalizeURL("https://example.com/path/")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Normalized URL: %s\n", url)
    }
}
