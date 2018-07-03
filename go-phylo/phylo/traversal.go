package phylo

// PreOrder traverses the tree in pre-order and calls fn for each node.
func PreOrder(n *Node, fn func(*Node)) {
    if n == nil || fn == nil {
        return
    }
    fn(n)
    for _, c := range n.Children {
        PreOrder(c, fn)
    }
}

// PostOrder traverses the tree in post-order.
func PostOrder(n *Node, fn func(*Node)) {
    if n == nil || fn == nil {
        return
    }
    for _, c := range n.Children {
        PostOrder(c, fn)
    }
    fn(n)
}

// CollectLeaves returns a slice with pointers to leaf nodes (left-to-right order).
func CollectLeaves(n *Node) []*Node {
    leaves := make([]*Node, 0)
    PreOrder(n, func(x *Node) {
        if x.IsLeaf() {
            leaves = append(leaves, x)
        }
    })
    return leaves
}
