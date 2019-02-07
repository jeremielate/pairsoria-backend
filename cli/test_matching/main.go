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

	var data struct {
		Patients []*match.Patient `json:"patients"`
	}
	data.Patients = make([]*match.Patient, 0, 1000)
	if err := dec.Decode(&data); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	couples := match.MatchPatients(data.Patients)

	out := bufio.NewWriter(os.Stdout)
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	if err := enc.Encode(couples); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	out.Flush()
}
