package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_ "embed"

	ivd "github.com/gopios/is-valid-domain"
)

//go:embed ../../public_suffix_list.dat
var pslData string

func main() {
	batchFlag := flag.String("batch", "", "File containing domains")
	flag.Parse()

	args := flag.Args()

	validator := ivd.New()

	if err := validator.LoadFromReader(strings.NewReader(pslData)); err != nil {
		log.Fatalf("Failed to load embedded PSL: %v", err)
	}

	fmt.Printf("Loaded %d suffixes\n", validator.SuffixCount())

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
	status := map[ivd.ValidationResult]string{
		ivd.Invalid:   "INVALID",
		ivd.Valid:     "VALID",
		ivd.Subdomain: "SUBDOMAIN",
	}[result]

	fmt.Printf("ivd %s - %d (%s)\n", domain, int(result), status)
}
