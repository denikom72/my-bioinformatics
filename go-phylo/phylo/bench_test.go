        package phylo

        import (
	"testing"
)

        func BenchmarkParseNewick(b *testing.B) {
            s := "((A:0.1,B:0.2)Inner:0.3,C:0.4)Root:0.0;"
            for i := 0; i < b.N; i++ {
                _, err := ParseNewick(s)
                if err != nil {
                    b.Fatalf("parse error: %v", err)
                }
            }
        }

        func BenchmarkPreOrder(b *testing.B) {
            // build a moderate tree
            root := NewNode("root", 0)
            for i := 0; i < 100; i++ {
                c := NewNode("n", 0)
                for j := 0; j < 10; j++ {
                    leaf := NewNode("leaf", 0)
                    c.AddChild(leaf)
                }
                root.AddChild(c)
            }
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                PreOrder(root, func(n *Node) {
                    _ = n.Name
                })
            }
        }
