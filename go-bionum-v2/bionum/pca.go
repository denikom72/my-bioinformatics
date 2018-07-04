package bionum

import (
    "gonum.org/v1/gonum/mat"
)

// PCA performs Principal Component Analysis on the given Matrix.
// Returns scores (projected rows) and loadings (principal axes)
func PCA(m *Matrix, nComponents int) (*Matrix, *Matrix) {
    r, c := m.R, m.C
    data := make([]float64, r*c)
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            data[i*c+j] = m.D[i][j]
        }
    }
    X := mat.NewDense(r, c, data)
    var svd mat.SVD
    ok := svd.Factorize(X, mat.SVDThin)
    if !ok {
        panic("SVD failed")
    }
    var U, V mat.Dense
    svd.UTo(&U)
    svd.VTo(&V)
    S := svd.Values(nil)

    // reduce components
    if nComponents > c { nComponents = c }
    scoresData := make([][]float64, r)
    for i := 0; i < r; i++ {
        scoresData[i] = make([]float64, nComponents)
        for j := 0; j < nComponents; j++ {
            scoresData[i][j] = U.At(i,j) * S[j]
        }
    }
    loadingsData := make([][]float64, c)
    for i := 0; i < c; i++ {
        loadingsData[i] = make([]float64, nComponents)
        for j := 0; j < nComponents; j++ {
            loadingsData[i][j] = V.At(i,j)
        }
    }
    return NewMatrix(scoresData), NewMatrix(loadingsData)
}
