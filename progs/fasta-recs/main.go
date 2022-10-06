package main

import (
	"fmt"
	"os"

	"birc.au.dk/gsa/fasta"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s fasta\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer f.Close()

	err = fasta.MapFasta(f, func(name, seq string) {
		fmt.Printf("%s\t%s\n", name, seq)
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
