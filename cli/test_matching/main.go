package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"pairsoria.com/server/match"
)

func main() {
	dec := json.NewDecoder(bufio.NewReader(os.Stdin))

	patients := make([]*match.Patient, 0, 1000)
	if err := dec.Decode(&patients); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	couples := match.MatchPatients(patients)

	out := bufio.NewWriter(os.Stdout)
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	if err := enc.Encode(couples); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	out.Flush()
}
