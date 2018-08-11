# go-motif

`go-motif` is a Go library for **motif discovery and matching** in DNA/RNA sequences.

## What is this for?

This library helps bioinformaticians detect regulatory sequence motifs in genomic data efficiently,
using Go-native parallel computation.

**Key terms (explained):**
- **Motif** (short recurring nucleotide pattern that may have biological function)
- **PWM** (Position Weight Matrix — matrix representing nucleotide probabilities at each position)
- **Regex-like pattern** (string pattern matching using nucleotide codes A,C,G,T,N)
- **p-value** (statistical significance of motif match)
- **Parallel scanning** (scanning sequences concurrently using Go routines for speed)

## Features
- PWM-based motif scoring
- Regex-like motif matching
- Parallel scanning of large sequence sets
- Simple API for adding motifs and scanning sequences
- Unit tests and benchmarks
- Makefile and GitHub Actions for CI

## Usage Example
```go
import (
    "fmt"
    "github.com/yourname/go-motif/motif"
)

func main() {
    seqs := []string{"ATGCGTAC", "CGTACGTA", "GCTAGCTA"}

    // define a simple PWM (2 positions example)
    pwm := motif.NewPWM([][]float64{
        {0.9,0.1,0.0,0.0}, // position 1: A,C,G,T
        {0.0,0.0,0.5,0.5}, // position 2
    })

    // Scan sequences in parallel
    results := motif.ScanSequences(seqs, pwm)

    for i, r := range results {
        fmt.Println("Sequence", i, "best score:", r.Score, "position:", r.Pos)
    }

    // Regex-like motif match
    matches := motif.MatchRegex(seqs, "ACG")
    fmt.Println("Regex matches:", matches)
}
```

## Installation
```bash
go get github.com/yourname/go-motif
```

## Run tests & benchmarks
```bash
make test
make bench
```
