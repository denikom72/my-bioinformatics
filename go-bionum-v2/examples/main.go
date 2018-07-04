package main

import (
    "fmt"
    "log"

    "github.com/yourname/go-bionum/bionum"
)

func main() {
    fmt.Println("go-bionum example")
    m := bionum.NewMatrix([][]float64{
        {100, 200, 50, 0},
        {120, 180, 60, 2},
        {90,  210, 40, 1},
    })

    tpm := bionum.TPMNormalize(m)
    tpm.LogTransform(1.0)
    fmt.Println("TPM log-transformed (row 0):", tpm.RowCopy(0))

    means := tpm.ColMeans()
    fmt.Println("Column means:", means)

    d := m.RowDistanceMatrix()
    fmt.Println("Distance matrix [0][1]:", d.At(0,1))

    corr := m.RowPearsonCorrelation()
    fmt.Println("Pearson corr [0][1]:", corr.At(0,1))

    // write example of reading CSV (skipped if file missing)
    if _, err := bionum.ReadCSV("data/example.csv", false); err != nil {
        log.Println("(no example.csv found, skipping CSV read)")
    }
}
