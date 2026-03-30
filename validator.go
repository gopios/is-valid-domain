package ivd

func New() *Validator {
	return &Validator{
		exact:     make(map[string]struct{}),
		wildcard:  make(map[string]struct{}),
		exception: make(map[string]struct{}),
	}
}
