package main

import (
    "fmt"
    "github.com/yourname/go-motif/motif"
)

func main() {
    seqs := []string{"ATGCGTAC","CGTACGTA","GCTAGCTA"}
    pwm := motif.NewPWM([][]float64{{0.9,0.1,0,0},{0,0,0.5,0.5}})
    results := motif.ScanSequences(seqs,pwm)
    for i,r := range results { fmt.Println("Seq",i,"best score:",r.Score,"pos:",r.Pos) }
    matches := motif.MatchRegex(seqs,"ACG")
    fmt.Println("Regex matches:",matches)
}
