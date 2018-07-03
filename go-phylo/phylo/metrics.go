package phylo

import (
    "errors"
)

// Ancestors returns the list of ancestors from node up to root (inclusive), root last.
func Ancestors(n *Node) []*Node {
    res := make([]*Node, 0)
    cur := n
    for cur != nil {
        res = append(res, cur)
        cur = cur.Parent
    }
    return res
}

// MRCA returns the most recent common ancestor of nodes a and b (or nil).
func MRCA(a, b *Node) *Node {
    if a == nil || b == nil {
        return nil
    }
    aAnc := Ancestors(a)
    bAnc := Ancestors(b)
    set := make(map[*Node]struct{})
    for _, x := range aAnc {
        set[x] = struct{}{}
    }
    // walk b's ancestors until a match
    for _, x := range bAnc {
        if _, ok := set[x]; ok {
            return x
        }
    }
    return nil
}

// PathLengthBetween returns the total branch length between two nodes by summing branch lengths
// from a up to MRCA and from b up to MRCA.
func PathLengthBetween(a, b *Node) (float64, error) {
    if a == nil || b == nil {
        return 0, errors.New("nil node")
    }
    m := MRCA(a, b)
    if m == nil {
        return 0, errors.New("no common ancestor")
    }
    var sum float64
    // from a up to but excluding m
    cur := a
    for cur != m {
        sum += cur.BranchLength
        cur = cur.Parent
        if cur == nil { // shouldn't happen
            return 0, errors.New("broken parent chain")
        }
    }
    // from b up to but excluding m
    cur = b
    for cur != m {
        sum += cur.BranchLength
        cur = cur.Parent
        if cur == nil {
            return 0, errors.New("broken parent chain")
        }
    }
    return sum, nil
}
