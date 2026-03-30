package ivd

import (
	"bufio"
	"io"
	"os"
	"strings"
)

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
