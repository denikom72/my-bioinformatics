package bionum

import (
    "encoding/csv"
    "os"
    "strconv"
)

// ReadCSV reads a numeric CSV into a Matrix. Assumes header may be present; pass hasHeader=true to skip first row.
func ReadCSV(path string, hasHeader bool) (*Matrix, error) {
    f, err := os.Open(path)
    if err != nil { return nil, err }
    defer f.Close()
    r := csv.NewReader(f)
    records, err := r.ReadAll()
    if err != nil { return nil, err }
    start := 0
    if hasHeader { start = 1 }
    rows := len(records) - start
    if rows <= 0 { return Zeros(0,0), nil }
    cols := len(records[start])
    data := make([][]float64, rows)
    for i := 0; i < rows; i++ {
        data[i] = make([]float64, cols)
        for j := 0; j < cols; j++ {
            v, _ := strconv.ParseFloat(records[i+start][j], 64)
            data[i][j] = v
        }
    }
    return NewMatrix(data), nil
}
