package phylo

import "testing"

func TestBasicTreeCounts(t *testing.T) {
    root := NewNode("root", 0)
    a := NewNode("A", 0.1)
    b := NewNode("B", 0.2)
    c := NewNode("C", 0.3)
    root.AddChild(a)
    root.AddChild(b)
    b.AddChild(c)

    if root.CountNodes() != 4 {
        t.Fatalf("expected 4 nodes, got %d", root.CountNodes())
    }
    if root.Depth() != 2 {
        t.Fatalf("expected depth 2, got %d", root.Depth())
    }
    if got := root.FindByName("C"); got == nil || got.Name != "C" {
        t.Fatalf("FindByName failed to find C")
    }
}

func TestMRCAAndDistance(t *testing.T) {
    // tree: (A:0.1,B:0.2)Inner:0.3;
    inner := NewNode("Inner", 0.3)
    a := NewNode("A", 0.1)
    b := NewNode("B", 0.2)
    inner.AddChild(a)
    inner.AddChild(b)
    root := NewNode("Root", 0.0)
    root.AddChild(inner)

    m := MRCA(a, b)
    if m == nil || m.Name != "Inner" {
        t.Fatalf("expected MRCA Inner, got %v", m)
    }
    dist, err := PathLengthBetween(a, b)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    // path: A (0.1) + B (0.2) = 0.3 (since Inner's branch length is to its parent Root)
    if dist <= 0 || (dist < 0.2999999 || dist > 0.3000001) {
        t.Fatalf("unexpected path length: %f", dist)
    }
}
