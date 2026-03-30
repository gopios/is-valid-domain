# is-valid-domain

A Go library and CLI tool for validating domain names against the Public Suffix List (PSL).

## Overview

This package provides domain validation functionality that determines whether a domain is:
- **VALID** - A registered domain
- **SUBDOMAIN** - A subdomain of a registered domain  
- **INVALID** - Not a valid domain

## Installation

### As a Go package

```bash
go get github.com/gopios/is-valid-domain
```

### As a CLI tool

```bash
go install github.com/gopios/is-valid-domain/cmd/ivd@latest
```

## Usage

### Library Usage

```go
package main

import (
    "fmt"
    "strings"
    ivd "github.com/gopios/is-valid-domain"
)

func main() {
    validator := ivd.New()
    
    // Load from embedded PSL data
    if err := validator.LoadFromReader(strings.NewReader(pslData)); err != nil {
        panic(err)
    }
    
    // Validate single domain
    result := validator.Validate("example.com")
    fmt.Printf("example.com: %v\n", result) // Output: example.com: VALID
    
    // Validate multiple domains
    domains := []string{"example.com", "sub.example.com", "invalid..com"}
    results := validator.ValidateBatch(domains)
    
    for domain, result := range results {
        fmt.Printf("%s: %v\n", domain, result)
    }
}
```

### CLI Usage

#### Single Domain Validation

```bash
ivd example.com
# Output: ivd example.com - 1 (VALID)

ivd sub.example.com  
# Output: ivd sub.example.com - 2 (SUBDOMAIN)

ivd invalid..com
# Output: ivd invalid..com - 0 (INVALID)
```

#### Batch Validation

Create a file with domains (one per line):

```
domains.txt
example.com
sub.example.com
test.org
invalid..com
```

Run batch validation:

```bash
ivd -batch domains.txt
# Output:
# ivd example.com - 1 (VALID)
# ivd sub.example.com - 2 (SUBDOMAIN)
# ivd test.org - 1 (VALID)
# ivd invalid..com - 0 (INVALID)
```

## API Reference

### Types

```go
type ValidationResult int

const (
    Invalid   ValidationResult = 0
    Valid     ValidationResult = 1
    Subdomain ValidationResult = 2
)
```

### Methods

#### `New() *Validator`
Creates a new validator instance.

#### `LoadFromReader(r io.Reader) error`
Loads Public Suffix List data from an io.Reader.

#### `Validate(domain string) ValidationResult`
Validates a single domain name.

#### `ValidateBatch(domains []string) map[string]ValidationResult`
Validates multiple domains and returns a map of results.

#### `SuffixCount() int`
Returns the total number of loaded suffixes (exact + wildcard + exception).

## Validation Rules

The validator follows the Public Suffix List rules:

1. **Exception Rules**: Domains that are exceptions to wildcard rules
2. **Exact Match Rules**: Domains that exactly match PSL entries
3. **Wildcard Rules**: Domains that match wildcard patterns

A domain is considered:
- **VALID** if it matches a registered domain exactly
- **SUBDOMAIN** if it's a subdomain of a registered domain
- **INVALID** if it doesn't match any PSL rules or has invalid format

## Examples

| Domain | Result | Explanation |
|--------|--------|-------------|
| `example.com` | VALID | Registered domain |
| `sub.example.com` | SUBDOMAIN | Subdomain of registered domain |
| `co.uk` | VALID | PSL entry (exact match) |
| `example.co.uk` | SUBDOMAIN | Subdomain of PSL entry |
| `invalid..com` | INVALID | Invalid format (double dot) |
| `com` | INVALID | TLD only, not a registered domain |
| `example.invalidtld` | INVALID | Invalid TLD |

## Performance

- Fast validation using hash maps for O(1) lookups
- Efficient batch processing
- Minimal memory footprint
- Embedded PSL data for standalone operation


## Public Suffix List

This library uses the [Public Suffix List](https://publicsuffix.org/) maintained by Mozilla. The PSL is an inclusive list of domain suffixes that users can rely on for privacy and security purposes.