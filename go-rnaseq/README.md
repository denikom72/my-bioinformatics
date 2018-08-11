# go-rnaseq

`go-rnaseq` is a Go library for **RNA-Seq data processing and analysis**.

## What is this for?

Provides basic tools to normalize and analyze gene expression matrices,
compute TPM/RPKM values, and prepare simple DE-like (Differential Expression) analysis.

**Key terms (explained):**
- **TPM** (Transcripts Per Million — normalization of RNA-Seq counts for sequencing depth and gene length)
- **RPKM/FPKM** (Reads/Fragments Per Kilobase per Million — alternative normalization method)
- **DE** (Differential Expression — detecting genes with significant expression changes)
- **Fold-change** (ratio of expression between conditions)
- **Library size** (total number of reads per sample)
- **Log2 transform** (logarithmic scaling to reduce skew)

## Features
- TPM normalization
- Log2 transformation
- Simple DE-like fold-change computation
- Parallel processing for large matrices
- CSV input/output helpers
- Unit tests and benchmarks
- Makefile and GitHub Actions for CI

## Usage Example
```go
import (
    "fmt"
    "github.com/yourname/go-rnaseq/rnaseq"
)

func main() {
    counts := [][]float64{
        {100, 200, 50},
        {120, 180, 60},
        {90, 210, 40},
    }

    mat := rnaseq.NewMatrix(counts)

    // TPM normalization
    tpm := rnaseq.TPMNormalize(mat)

    // Log2 transform
    rnaseq.LogTransform(tpm, 1.0)

    // Compute fold-change between sample 0 and 1
    fc := rnaseq.FoldChange(tpm, 0, 1)
    fmt.Println("Fold-change sample0 vs sample1:", fc)
}
```

## Installation
```bash
go get github.com/yourname/go-rnaseq
```

## Run tests & benchmarks
```bash
make test
make bench
```
