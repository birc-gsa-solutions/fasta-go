package main

import (
	"bufio"
	"fmt"
	"os"

	"birc.au.dk/gsa/fasta"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s fasta [coordinates]\n", os.Args[0])
		os.Exit(1)
	}

	fastaFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	var fastaRecs map[string]string
	fastaRecs, err = fasta.LoadFasta(fastaFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fastaFile.Close()

	// Get the coordinates stream from either a file
	// or stdin
	var coordFile = os.Stdin
	if len(os.Args) == 3 && os.Args[2] != "-" {
		coordFile, err = os.Open(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		defer coordFile.Close()
	}

	// Then process each line in the coordinates stream
	var (
		name  string
		start int
		stop  int
	)
	scanner := bufio.NewScanner(coordFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "%s\t%d\t%d\n", &name, &start, &stop)
		fmt.Println(fastaRecs[name][start-1 : stop-1])
	}

}
