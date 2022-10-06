package fasta

import (
	"fmt"
	"io"
	"strings"
)

// MapFasta reads a fasta file from f and maps the function f
// over all records. f is called with the record name and the
// record sequence
func MapFasta(r io.Reader, f func(string, string)) error {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	records := strings.Split(string(bytes), ">")
	if len(records) == 0 {
		// empty, it shouldn't happen, but we will consider
		// it valid...
		return nil
	}

	if records[0] != "" {
		return fmt.Errorf("Expected an empty string before first header")
	}

	for i := 1; i < len(records); i++ {
		lines := strings.Split(records[i], "\n")
		header := strings.TrimSpace(lines[0])
		seq := strings.Join(lines[1:], "")
		f(header, seq)
	}

	return nil
}

// LoadFasta loads a fasta file into a map that maps
// from record names to sequences.
func LoadFasta(r io.Reader) (map[string]string, error) {
	m := map[string]string{}
	err := MapFasta(r, func(name, seq string) {
		m[name] = seq
	})

	if err != nil {
		return nil, err
	}

	return m, nil
}
