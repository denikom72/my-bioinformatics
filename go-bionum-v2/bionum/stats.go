        package bionum

        import (
	"math"
)

        // RowMeans computes mean of each row (returns slice length R)
        func (m *Matrix) RowMeans() []float64 {
            out := make([]float64, m.R)
            for i := 0; i < m.R; i++ {
                var s float64
                for j := 0; j < m.C; j++ {
                    s += m.D[i][j]
                }
                out[i] = s / float64(m.C)
            }
            return out
        }

        // ColMeans computes mean of each column (returns slice length C)
        func (m *Matrix) ColMeans() []float64 {
            out := make([]float64, m.C)
            for j := 0; j < m.C; j++ {
                var s float64
                for i := 0; i < m.R; i++ {
                    s += m.D[i][j]
                }
                out[j] = s / float64(m.R)
            }
            return out
        }

        func variance(xs []float64) float64 {
            if len(xs) == 0 { return 0 }
            mean := 0.0
            for _, v := range xs { mean += v }
            mean /= float64(len(xs))
            var s float64
            for _, v := range xs {
                d := v - mean
                s += d * d
            }
            return s / float64(len(xs))
        }

        func stddev(xs []float64) float64 { return math.Sqrt(variance(xs)) }

        // ColStd returns column standard deviations
        func (m *Matrix) ColStd() []float64 {
            out := make([]float64, m.C)
            for j := 0; j < m.C; j++ {
                col := m.ColCopy(j)
                out[j] = stddev(col)
            }
            return out
        }

        // ZScoreColumns performs z-score standardization column-wise in-place
        func (m *Matrix) ZScoreColumns() {
            means := m.ColMeans()
            sds := m.ColStd()
            for j := 0; j < m.C; j++ {
                sd := sds[j]
                if sd == 0 { sd = 1 }
                for i := 0; i < m.R; i++ {
                    m.D[i][j] = (m.D[i][j] - means[j]) / sd
                }
            }
        }

        // LogTransform applies log2(x + pseudocount) in-place
        func (m *Matrix) LogTransform(pseudocount float64) {
            for i := 0; i < m.R; i++ {
                for j := 0; j < m.C; j++ {
                    m.D[i][j] = math.Log2(m.D[i][j] + pseudocount)
                }
            }
        }

        // TPMNormalize assumes rows are samples and columns are transcripts (counts per transcript)
        // For each row: divide counts by total counts (per-row sum), multiply by 1e6
        func TPMNormalize(m *Matrix) *Matrix {
            out := Zeros(m.R, m.C)
            for i := 0; i < m.R; i++ {
                var sum float64
                for j := 0; j < m.C; j++ { sum += m.D[i][j] }
                if sum == 0 { sum = 1 }
                for j := 0; j < m.C; j++ { out.D[i][j] = m.D[i][j] / sum * 1e6 }
            }
            return out
        }

        // RowDistanceMatrix computes Euclidean distance between rows (samples)
        func (m *Matrix) RowDistanceMatrix() *Matrix {
            out := Zeros(m.R, m.R)
            for i := 0; i < m.R; i++ {
                for j := i; j < m.R; j++ {
                    var sum float64
                    for k := 0; k < m.C; k++ {
                        d := m.D[i][k] - m.D[j][k]
                        sum += d * d
                    }
                    dist := math.Sqrt(sum)
                    out.D[i][j] = dist
                    out.D[j][i] = dist
                }
            }
            return out
        }

        // PearsonCorrelationMatrix computes pairwise Pearson correlations between rows
        func (m *Matrix) RowPearsonCorrelation() *Matrix {
            out := Zeros(m.R, m.R)
            means := m.RowMeans()
            for i := 0; i < m.R; i++ {
                for j := i; j < m.R; j++ {
                    var num float64
                    var si, sj float64
                    for k := 0; k < m.C; k++ {
                        xi := m.D[i][k] - means[i]
                        xj := m.D[j][k] - means[j]
                        num += xi * xj
                        si += xi * xi
                        sj += xj * xj
                    }
                    denom := math.Sqrt(si * sj)
                    corr := 0.0
                    if denom != 0 { corr = num / denom }
                    out.D[i][j] = corr
                    out.D[j][i] = corr
                }
            }
            return out
        }
