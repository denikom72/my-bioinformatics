package bionum

import "errors"

// Matrix is a simple 2D numeric matrix (rows x cols).
type Matrix struct {
    R int
    C int
    D [][]float64
}

// NewMatrix creates a new Matrix; input must be rectangular.
func NewMatrix(data [][]float64) *Matrix {
    if len(data) == 0 {
        return &Matrix{R: 0, C: 0, D: [][]float64{}}
    }
    r := len(data)
    c := len(data[0])
    // ensure rectangular
    for _, row := range data {
        if len(row) != c {
            panic("rows have different lengths")
        }
    }
    return &Matrix{R: r, C: c, D: data}
}

func Zeros(r, c int) *Matrix {
    d := make([][]float64, r)
    for i := 0; i < r; i++ {
        d[i] = make([]float64, c)
    }
    return &Matrix{R: r, C: c, D: d}
}

func (m *Matrix) At(i, j int) float64 { return m.D[i][j] }
func (m *Matrix) Set(i, j int, v float64) { m.D[i][j] = v }

// Transpose returns a new matrix with rows and columns swapped.
func (m *Matrix) Transpose() *Matrix {
    t := Zeros(m.C, m.R)
    for i := 0; i < m.R; i++ {
        for j := 0; j < m.C; j++ {
            t.D[j][i] = m.D[i][j]
        }
    }
    return t
}

// Multiply performs matrix multiplication (this * other).
func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
    if m.C != other.R {
        return nil, errors.New("incompatible dimensions")
    }
    out := Zeros(m.R, other.C)
    for i := 0; i < m.R; i++ {
        for j := 0; j < other.C; j++ {
            var sum float64
            for k := 0; k < m.C; k++ {
                sum += m.D[i][k] * other.D[k][j]
            }
            out.D[i][j] = sum
        }
    }
    return out, nil
}

// RowCopy returns a copy of a row.
func (m *Matrix) RowCopy(i int) []float64 {
    c := make([]float64, m.C)
    copy(c, m.D[i])
    return c
}

// ColCopy returns a copy of a column.
func (m *Matrix) ColCopy(j int) []float64 {
    r := make([]float64, m.R)
    for i := 0; i < m.R; i++ {
        r[i] = m.D[i][j]
    }
    return r
}
