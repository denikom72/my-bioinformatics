# go-phylo (minimal, with CI & Makefile)

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
