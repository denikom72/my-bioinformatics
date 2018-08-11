package rnaseq

import "testing"

func TestTPMNormalize(t *testing.T) {
    mat := NewMatrix([][]float64{{100,200},{50,50}})
    tpm := TPMNormalize(mat)
    if len(tpm.Data) != 2 || len(tpm.Data[0]) != 2 { t.Fatalf("unexpected shape") }
}

func TestLogTransform(t *testing.T) {
    mat := NewMatrix([][]float64{{1,3},{4,5}})
    LogTransform(mat,1.0)
    if mat.Data[0][0] <= 0 { t.Fatalf("log-transform failed") }
}

func TestFoldChange(t *testing.T) {
    mat := NewMatrix([][]float64{{10,5},{20,10}})
    fc := FoldChange(mat,0,1)
    if fc[0] != 2 { t.Fatalf("unexpected fold-change") }
}
