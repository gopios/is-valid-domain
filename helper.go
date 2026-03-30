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

	root := strings.Join(parts[len(parts)-len(suffixParts)-1:], ".")

	if domain == root {
		return Valid
	}
	return Subdomain
}
