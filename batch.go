package ivd

func (v *Validator) ValidateBatch(domains []string) map[string]ValidationResult {
	results := make(map[string]ValidationResult, len(domains))

	for _, d := range domains {
		results[d] = v.Validate(d)
	}

	return results
}
