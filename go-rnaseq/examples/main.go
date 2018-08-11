package main

import (
    "fmt"
    "github.com/yourname/go-rnaseq/rnaseq"
)

func main() {
    counts := [][]float64{
        {100, 200, 50},
        {120, 180, 60},
        {90, 210, 40},
    }
    mat := rnaseq.NewMatrix(counts)
    tpm := rnaseq.TPMNormalize(mat)
    rnaseq.LogTransform(tpm,1.0)
    fc := rnaseq.FoldChange(tpm,0,1)
    fmt.Println("Fold-change sample0 vs sample1:",fc)
}
