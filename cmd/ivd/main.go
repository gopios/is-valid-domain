package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	ivd "github.com/gopios/is-valid-domain"
)

func main() {
	batchFlag := flag.String("batch", "", "File containing domains")
	flag.Parse()

	args := flag.Args()

	validator := ivd.NewWithPSL()

	// Batch mode
	if *batchFlag != "" {
		data, err := os.ReadFile(*batchFlag)
		if err != nil {
			log.Fatalf("Failed to read batch file: %v", err)
		}

		domains := strings.Split(string(data), "\n")
		results := validator.ValidateBatch(domains)

		for domain, result := range results {
			domain = strings.TrimSpace(domain)
			if domain != "" {
				printResult(domain, result)
			}
		}
		return
	}

	// CLI args
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, domain := range args {
		printResult(domain, validator.Validate(domain))
	}
}

func printResult(domain string, result ivd.ValidationResult) {
	fmt.Printf("%d\n", int(result))
}
