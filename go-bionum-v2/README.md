# go-bionum (v2)

`go-bionum` is a lightweight Go library providing **numerical and statistical utilities**
tailored for bioinformatics data (e.g. gene expression matrices, distance matrices).

## What is this for?

This library helps researchers and engineers process and analyze biological datasets
with efficient, idiomatic Go code. Typical tasks include normalization of expression
matrices, computation of summary statistics, distance matrices, dimensionality reduction (PCA/SVD),
and basic data transformations used in downstream analyses (e.g., clustering, visualization).

**Key terms (explained):**
- **Matrix** (2D array of numbers; rows often represent samples, columns represent features/genes)
- **TPM** (Transcripts Per Million — normalization for RNA-Seq counts)
- **Z-score** (standardized value: (x - mean) / std; used to compare values across features)
- **Distance matrix** (pairwise distances between samples; used for clustering)
- **Pearson correlation** (measure of linear correlation between two vectors)
- **PCA** (Principal Component Analysis — projects high-dimensional data into lower dimensions maximizing variance)
- **SVD** (Singular Value Decomposition — factorizes a matrix into orthogonal components, foundation for PCA)

## Features
- Basic `Matrix` type with creation, transpose, multiply
- Row/column means, variances, standard deviation
- Column-wise TPM normalization and log2 transform
- Z-score normalization
- Euclidean distance matrix and Pearson correlation matrix
- Dimensionality reduction: PCA, SVD (using gonum or native)
- CSV input/output helpers
- Unit tests and benchmarks
- Makefile and GitHub Actions for CI

## Usage (examples)
```go
import "github.com/denikom72/go-bionum/bionum"

m := bionum.NewMatrix([][]float64{
    {100, 200, 50, 0},
    {120, 180, 60, 2},
    {90,  210, 40, 1},
})

// Normalize, log-transform, z-score
tpm := bionum.TPMNormalize(m)
bionum.LogTransform(tpm, 1.0)
tpm.ZScoreColumns()

// Compute distance and correlation
dmat := tpm.RowDistanceMatrix()
corr := tpm.RowPearsonCorrelation()

// PCA example (reduce to 2 components)
scores, loadings := bionum.PCA(tpm, 2)
fmt.Println("PCA scores (first row):", scores.RowCopy(0))
fmt.Println("PCA loadings (first component):", loadings.ColCopy(0))
```

## Installation
```bash
go get github.com/denikom72/go-bionum
```

## Run tests & benchmarks
```bash
make test
make bench
```

## Notes
Gonum is optionally used internally for SVD for better performance but all functions
work with pure Go code fallback.
