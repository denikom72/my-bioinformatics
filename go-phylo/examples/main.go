package main

import (
    "fmt"
    "log"

    "github.com/yourname/go-phylo/phylo"
)

func main() {
    fmt.Println("=== RF-Distanz Beispiel ===")
    t1 := phylo.NewNode("Root", 0)
    a := phylo.NewNode("A", 0)
    b := phylo.NewNode("B", 0)
    t1.AddChild(a)
    t1.AddChild(b)
    t2 := phylo.NewNode("Root", 0)
    b2 := phylo.NewNode("B", 0)
    a2 := phylo.NewNode("A", 0)
    t2.AddChild(b2)
    t2.AddChild(a2)
    dist, err := phylo.RobinsonFouldsDistance(t1, t2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("RF-Distanz t1 vs t2:", dist)

    fmt.Println("=== Neighbor-Joining Beispiel ===")
    labels := []string{"A","B","C","D","E"}
    matrix := [][]float64{
        {0,5,9,9,8},
        {5,0,10,10,9},
        {9,10,0,8,7},
        {9,10,8,0,3},
        {8,9,7,3,0},
    }
    njTree, err := phylo.NeighborJoining(matrix, labels)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("NJ-Baum Wurzel:", njTree)

    fmt.Println("=== UPGMA Beispiel ===")
    upgmaTree, err := phylo.UPGMA(matrix, labels)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("UPGMA-Baum Wurzel:", upgmaTree)
}
