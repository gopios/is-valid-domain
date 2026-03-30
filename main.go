package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <domain>\n", os.Args[0])
		os.Exit(1)
	}

	domain := strings.TrimSpace(os.Args[1])
	if domain == "" {
		fmt.Println("invalid")
		os.Exit(1)
	}

	// Get the effective TLD+1 (registrable domain)
	etld1, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		fmt.Println("invalid")
		os.Exit(1)
	}

	// If the domain is exactly the effective TLD+1, it's a valid domain
	// If it's longer, it's a subdomain
	if domain == etld1 {
		fmt.Println("valid")
	} else {
		fmt.Println("sub")
	}
}
