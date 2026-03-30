package ivd

import (
	"bufio"
	_ "embed"
	"io"
	"os"
	"strings"
)

//go:embed public_suffix_list.dat
var pslData string

func (v *Validator) LoadFromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return v.load(f)
}

func (v *Validator) LoadFromReader(r io.Reader) error {
	return v.load(r)
}

// NewWithPSL creates a new Validator and automatically loads the embedded Public Suffix List
func NewWithPSL() *Validator {
	validator := New()
	if err := validator.LoadFromReader(strings.NewReader(pslData)); err != nil {
		// In practice this should never happen since the data is embedded
		panic("failed to load embedded PSL data: " + err.Error())
	}
	return validator
}

func (v *Validator) load(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "!"):
			v.exception[line[1:]] = struct{}{}

		case strings.HasPrefix(line, "*."):
			v.wildcard[line[2:]] = struct{}{}

		default:
			v.exact[line] = struct{}{}
		}
	}

	return scanner.Err()
}
