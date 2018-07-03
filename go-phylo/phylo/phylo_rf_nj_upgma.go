package phylo

import "testing"

func TestRFDistanceSimple(t *testing.T) {
    // simple tree1: (A,B)Root;
    n1 := NewNode("Root", 0)
    a := NewNode("A", 0)
    b := NewNode("B", 0)
    n1.AddChild(a)
    n1.AddChild(b)

    // simple tree2: (B,A)Root;
    n2 := NewNode("Root", 0)
    b2 := NewNode("B", 0)
    a2 := NewNode("A", 0)
    n2.AddChild(b2)
    n2.AddChild(a2)

    dist, err := RobinsonFouldsDistance(n1, n2)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if dist != 0 {
        t.Fatalf("expected 0 RF distance, got %d", dist)
    }
}

func TestUPGMAAndNJ(t *testing.T) {
    dist := [][]float64{
        {0, 5, 9, 9, 8},
        {5, 0, 10, 10, 9},
        {9, 10, 0, 8, 7},
        {9, 10, 8, 0, 3},
        {8, 9, 7, 3, 0},
    }
    labels := []string{"A","B","C","D","E"}
    upgmaTree, err := UPGMA(dist, labels)
    if err != nil || upgmaTree == nil {
        t.Fatalf("UPGMA failed: %v", err)
    }
    njTree, err := NeighborJoining(dist, labels)
    if err != nil || njTree == nil {
        t.Fatalf("NJ failed: %v", err)
    }
}
