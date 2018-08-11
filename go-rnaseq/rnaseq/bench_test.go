package rnaseq

import "testing"

func BenchmarkTPMNormalize(b *testing.B) {
    mat := make([][]float64,1000)
    for i:=0;i<1000;i++ { mat[i] = []float64{100,200,50,75} }
    m := NewMatrix(mat)
    b.ResetTimer()
    for i:=0;i<b.N;i++ { _ = TPMNormalize(m) }
}
