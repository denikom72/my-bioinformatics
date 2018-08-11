package motif

import "testing"

func BenchmarkScanSequences(b *testing.B) {
    seqs := make([]string, 1000)
    for i := range seqs { seqs[i] = "ACGTACGTACGT" }
    pwm := NewPWM([][]float64{{0.25,0.25,0.25,0.25}})
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = ScanSequences(seqs,pwm)
    }
}
