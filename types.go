package ivd

type ValidationResult int

const (
	Invalid ValidationResult = iota
	Valid
	Subdomain
)

type Validator struct {
	exact     map[string]struct{}
	wildcard  map[string]struct{}
	exception map[string]struct{}
}
