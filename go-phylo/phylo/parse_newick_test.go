package phylo

import "testing"

func TestParseNewickSimple(t *testing.T) {
    s := "((A:0.1,B:0.2)Inner:0.3,C:0.4)Root:0.0;"
    n, err := ParseNewick(s)
    if err != nil {
        t.Fatalf("parse error: %v", err)
    }
    if n == nil {
        t.Fatalf("parsed nil")
    }
    if n.FindByName("A") == nil || n.FindByName("B") == nil || n.FindByName("C") == nil {
        t.Fatalf("expected leaves A, B, C to be present")
    }
    // round-trip
    w := WriteNewick(n)
    if w == "" {
        t.Fatalf("write produced empty string")
    }
}
