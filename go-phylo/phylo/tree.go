package phylo

import (
    "errors"
    "fmt"
)

// Node represents a node in a phylogenetic tree.
type Node struct {
    Name         string
    BranchLength float64
    Parent       *Node
    Children     []*Node
}

// NewNode creates a new node with an optional name and branch length.
func NewNode(name string, branchLength float64) *Node {
    return &Node{Name: name, BranchLength: branchLength}
}

// AddChild appends a child to the node and sets its parent.
func (n *Node) AddChild(child *Node) {
    child.Parent = n
    n.Children = append(n.Children, child)
}

// IsLeaf returns true when the node has no children.
func (n *Node) IsLeaf() bool {
    return len(n.Children) == 0
}

// CountNodes returns total nodes in the subtree rooted at n (including n).
func (n *Node) CountNodes() int {
    if n == nil {
        return 0
    }
    count := 1
    for _, c := range n.Children {
        count += c.CountNodes()
    }
    return count
}

// Depth returns the maximum depth of the subtree (number of edges to deepest leaf).
func (n *Node) Depth() int {
    if n == nil {
        return 0
    }
    if n.IsLeaf() {
        return 0
    }
    max := 0
    for _, c := range n.Children {
        d := c.Depth() + 1
        if d > max {
            max = d
        }
    }
    return max
}

// FindByName searches for a node by name (first match, DFS).
func (n *Node) FindByName(name string) *Node {
    if n == nil {
        return nil
    }
    if n.Name == name {
        return n
    }
    for _, c := range n.Children {
        if found := c.FindByName(name); found != nil {
            return found
        }
    }
    return nil
}

// Reroot makes newRoot the root of the tree and adjusts parent/children pointers.
// This operation is destructive (changes pointers) and returns an error if newRoot is nil.
func (n *Node) Reroot(newRoot *Node) error {
    if newRoot == nil {
        return errors.New("newRoot is nil")
    }
    // Walk up from newRoot to current root, reversing edges.
    cur := newRoot
    var prev *Node = nil
    for cur != nil {
        parent := cur.Parent
        // remove cur from parent's children (if parent exists)
        if parent != nil {
            newChildren := make([]*Node, 0, len(parent.Children))
            for _, ch := range parent.Children {
                if ch != cur {
                    newChildren = append(newChildren, ch)
                }
            }
            parent.Children = newChildren
        }
        // reverse pointer
        cur.Parent = prev
        if prev != nil {
            prev.Children = append(prev.Children, cur)
        }
        // continue
        prev = cur
        cur = parent
    }
    return nil
}

// String returns a simple representation of the node (name or leaf count).
func (n *Node) String() string {
    if n == nil {
        return "<nil>"
    }
    if n.Name != "" {
        return n.Name
    }
    return fmt.Sprintf("<node:%d>", n.CountNodes())
}
