# go-phylo (minimal, with CI & Makefile)

go-phylo is a lightweight, high-performance library for phylogenetic tree analysis and manipulation written in Go.
A phylogenetic tree (i.e., a branching diagram that represents the evolutionary relationships among species, genes, or other biological entities) helps researchers understand how organisms are related and how they evolved over time.

The library is designed for computational biologists, bioinformaticians, and developers who need fast, memory-safe tree operations without relying on heavy Python or R frameworks.

It enables you to:

    Parse, build, and traverse phylogenetic trees in Newick format

    Compute distances, common ancestors, and path lengths

    Perform tree comparisons using the Robinson–Foulds (RF) distance

    Reconstruct trees from distance matrices using Neighbor-Joining (NJ) and UPGMA algorithms

    Integrate easily into pipelines or web services thanks to Go’s concurrency and efficiency

go-phylo aims to bring the speed and simplicity of Go into the bioinformatics world, offering a foundation for building scalable tools for evolutionary analysis, clustering, and genomics data processing.


## Features
- Tree structures (Node, Tree)
- Newick parse/write
- Traversal: PreOrder/PostOrder
- MRCA and path distance
- Unit tests and benchmarks
- Continuous Integration with GitHub Actions
- Makefile shortcuts

## USAGE
1. Clone repo and adjust `go.mod` module path.
2. Run example: `go run ./examples`.
3. Run tests: `make test` or `go test ./...`.
4. Run benchmarks: `make bench` or `go test -bench=. ./...`
5. Build binary: `make build`
