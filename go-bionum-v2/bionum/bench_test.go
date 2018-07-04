package bionum

import "testing"

func BenchmarkRowDistance(b *testing.B) {
    m := Zeros(200, 100)
    for i := 0; i < m.R; i++ {
        for j := 0; j < m.C; j++ {
            m.D[i][j] = float64(i+j)
        }
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = m.RowDistanceMatrix()
    }
}
