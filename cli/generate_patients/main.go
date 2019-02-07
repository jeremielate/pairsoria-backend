package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"pairsoria.com/server/match"
)

func usageAndExit() {
	fmt.Fprintf(os.Stderr, "usage: %v patients", os.Args[0])
	os.Exit(1)

}

func generateFakePatients(count int) []match.Patient {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	patients := make([]match.Patient, count)
	for i := 0; i < count; i++ {
		patients[i] = match.Patient{
			Id:     uuid.New(),
			Age:    rnd.Int() % 100,
			Points: rnd.Uint64() % 100000,
			Genre:  match.Genre(rnd.Int() % 3),
		}
	}
	return patients
}

func main() {
	if len(os.Args) != 2 {
		usageAndExit()
	}
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		usageAndExit()
	}

	var data struct {
		Patients []match.Patient `json:"patients"`
	}
	data.Patients = generateFakePatients(count)

	out := bufio.NewWriter(os.Stdout)
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	out.Flush()
}
