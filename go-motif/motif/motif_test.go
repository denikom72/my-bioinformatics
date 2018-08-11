package motif

import "testing"

func TestPWMScore(t *testing.T) {
    pwm := NewPWM([][]float64{
        {1,0,0,0},
        {0,1,0,0},
    })
    pos, score := pwm.ScoreSequence("AC")
    if pos != 0 || score != 2 { t.Fatalf("unexpected score: %v", score) }
}

func TestScanSequences(t *testing.T) {
    pwm := NewPWM([][]float64{{1,0,0,0}})
    seqs := []string{"A","C","A"}
    res := ScanSequences(seqs,pwm)
    if len(res) != 3 { t.Fatalf("unexpected length") }
}

func TestMatchRegex(t *testing.T) {
    seqs := []string{"ACG","GTA"}
    matches := MatchRegex(seqs,"AC")
    if len(matches[0]) != 1 || matches[0][0] != 0 { t.Fatalf("unexpected match") }
}
