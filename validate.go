package ivd

import "strings"

func (v *Validator) Validate(domain string) ValidationResult {
	domain = strings.ToLower(strings.TrimSpace(domain))

	if domain == "" || strings.Contains(domain, "..") {
		return Invalid
	}

	parts := strings.Split(domain, ".")
	n := len(parts)

	if n < 2 {
		return Invalid
	}

	// 1. Exception
	for i := 0; i < n; i++ {
		s := strings.Join(parts[i:], ".")
		if _, ok := v.exception[s]; ok {
			suffixParts := strings.Split(s, ".")
			suffix := strings.Join(suffixParts[1:], ".")
			return v.checkRoot(domain, suffix)
		}
	}

	// 2. Exact
	for i := 0; i < n; i++ {
		s := strings.Join(parts[i:], ".")
		if _, ok := v.exact[s]; ok {
			return v.checkRoot(domain, s)
		}
	}

	// 3. Wildcard
	for i := 1; i < n; i++ {
		s := strings.Join(parts[i:], ".")
		if _, ok := v.wildcard[s]; ok {
			return v.checkRoot(domain, s)
		}
	}

	return Invalid
}
