package rnaseq

import (
    "math"
)

type Matrix struct {
    Data [][]float64
    R, C int
}

func NewMatrix(data [][]float64) *Matrix {
    r := len(data)
    c := 0
    if r > 0 { c = len(data[0]) }
    return &Matrix{Data: data, R: r, C: c}
}

// TPMNormalize normalizes counts to Transcripts Per Million
func TPMNormalize(m *Matrix) *Matrix {
    r, c := m.R, m.C
    sums := make([]float64, r)
    for i := 0; i < r; i++ {
        s := 0.0
        for j := 0; j < c; j++ { s += m.Data[i][j] }
        sums[i] = s
    }
    out := make([][]float64, r)
    for i := 0; i < r; i++ {
        out[i] = make([]float64, c)
        for j := 0; j < c; j++ {
            if sums[i] != 0 {
                out[i][j] = m.Data[i][j] / sums[i] * 1e6
            } else { out[i][j] = 0 }
        }
    }
    return NewMatrix(out)
}

func LogTransform(m *Matrix, pseudo float64) {
    for i := 0; i < m.R; i++ {
        for j := 0; j < m.C; j++ {
            m.Data[i][j] = math.Log2(m.Data[i][j] + pseudo)
        }
    }
}

func FoldChange(m *Matrix, col1, col2 int) []float64 {
    fc := make([]float64, m.R)
    for i := 0; i < m.R; i++ {
        if m.Data[i][col2] != 0 {
            fc[i] = m.Data[i][col1] / m.Data[i][col2]
        } else {
            fc[i] = 0
        }
    }
    return fc
}
