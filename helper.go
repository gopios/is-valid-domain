package ivd

import "strings"

func (v *Validator) SuffixCount() int {
	return len(v.exact) + len(v.wildcard) + len(v.exception)
}

func (v *Validator) checkRoot(domain, suffix string) ValidationResult {
	parts := strings.Split(domain, ".")
	suffixParts := strings.Split(suffix, ".")

	if len(parts) <= len(suffixParts) {
		return Invalid
	}

	// If domain has exactly one more part than suffix, it's a registered domain
	if len(parts) == len(suffixParts)+1 {
		return Valid
	}

	// Otherwise it's a subdomain
	return Subdomain
}
