package motif

type PWM struct {
    Matrix [][]float64 // positions x nucleotides (A,C,G,T)
}

func NewPWM(matrix [][]float64) *PWM {
    return &PWM{Matrix: matrix}
}

// ScoreSequence returns best PWM score in a sequence
func (p *PWM) ScoreSequence(seq string) (pos int, score float64) {
    n := len(seq)
    m := len(p.Matrix)
    best := -1.0
    bestPos := 0
    for i := 0; i <= n-m; i++ {
        s := 0.0
        for j := 0; j < m; j++ {
            nt := seq[i+j]
            idx := ntToIndex(nt)
            s += p.Matrix[j][idx]
        }
        if s > best {
            best = s
            bestPos = i
        }
    }
    return bestPos, best
}

func ntToIndex(nt byte) int {
    switch nt {
    case 'A': return 0
    case 'C': return 1
    case 'G': return 2
    case 'T': return 3
    default: return 0
    }
}

type Result struct {
    Pos int
    Score float64
}

// ScanSequences scans multiple sequences in parallel using PWM
func ScanSequences(seqs []string, pwm *PWM) []Result {
    res := make([]Result, len(seqs))
    done := make(chan struct{}, len(seqs))
    for i, s := range seqs {
        go func(i int, seq string) {
            pos, score := pwm.ScoreSequence(seq)
            res[i] = Result{Pos: pos, Score: score}
            done <- struct{}{}
        }(i, s)
    }
    for i := 0; i < len(seqs); i++ { <-done }
    return res
}

// MatchRegex performs simple substring matching (regex-like) on sequences
func MatchRegex(seqs []string, motif string) map[int][]int {
    out := make(map[int][]int)
    for i, s := range seqs {
        positions := []int{}
        n := len(s)
        m := len(motif)
        for j := 0; j <= n-m; j++ {
            match := true
            for k := 0; k < m; k++ {
                if s[j+k] != motif[k] { match = false; break }
            }
            if match { positions = append(positions, j) }
        }
        out[i] = positions
    }
    return out
}
