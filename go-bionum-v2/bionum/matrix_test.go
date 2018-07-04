package bionum

import "testing"

func TestRowColMeans(t *testing.T) {
    m := NewMatrix([][]float64{{1,2,3},{4,5,6}})
    rm := m.RowMeans()
    if len(rm) != 2 || rm[0] != 2 { t.Fatalf("unexpected row means: %v", rm) }
    cm := m.ColMeans()
    if len(cm) != 3 || cm[0] != 2.5 { t.Fatalf("unexpected col means: %v", cm) }
}

func TestTPMAndLog(t *testing.T) {
    m := NewMatrix([][]float64{{100,200},{50,50}})
    tpm := TPMNormalize(m)
    if tpm.D[0][0] <= 0 { t.Fatalf("tpm zero") }
    tpm.LogTransform(1.0)
    // ensure transformed values exist
    if tpm.D[0][0] == tpm.D[0][1] { t.Fatalf("unexpected log transform") }
}

func TestDistanceAndCorrelation(t *testing.T) {
    m := NewMatrix([][]float64{{1,2,3},{1,2,3},{2,2,2}})
    d := m.RowDistanceMatrix()
    if d.D[0][1] != 0 { t.Fatalf("expected zero distance between identical rows") }
    corr := m.RowPearsonCorrelation()
    if corr.D[0][1] < 0.999 { t.Fatalf("expected high corr: %v", corr.D[0][1]) }
}
