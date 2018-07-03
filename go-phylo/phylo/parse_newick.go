package phylo

import (
    "fmt"
    "strconv"
    "strings"
    "unicode"
)

// ParseNewick parses a Newick formatted string and returns the root node.
// Supports labels and branch lengths: (A:0.1,B:0.2)Root:0.0;
func ParseNewick(s string) (*Node, error) {
    orig := s
    s = strings.TrimSpace(s)
    if s == "" {
        return nil, fmt.Errorf("empty string")
    }
    // remove trailing semicolon if present
    if strings.HasSuffix(s, ";") {
        s = strings.TrimSuffix(s, ";")
    }
    p := &newickParser{input: s, pos: 0}
    node, err := p.parseSubtree()
    if err != nil {
        return nil, err
    }
    // ensure consumed whole string
    p.skipWhitespace()
    if p.pos < len(p.input) {
        return nil, fmt.Errorf("unexpected trailing: %q in %q", p.input[p.pos:], orig)
    }
    return node, nil
}

// WriteNewick writes a subtree rooted at n to a Newick string (no trailing semicolon).
func WriteNewick(n *Node) string {
    if n == nil {
        return ";"
    }
    var sb strings.Builder
    writeNode(&sb, n)
    return sb.String()
}

func writeNode(sb *strings.Builder, n *Node) {
    if n.IsLeaf() {
        if n.Name != "" {
            sb.WriteString(escapeLabel(n.Name))
        }
        if !isZero(n.BranchLength) {
            sb.WriteString(":" + formatFloat(n.BranchLength))
        }
        return
    }
    // internal node with children
    sb.WriteRune('(')
    for i, c := range n.Children {
        if i > 0 {
            sb.WriteRune(',')
        }
        writeNode(sb, c)
    }
    sb.WriteRune(')')
    if n.Name != "" {
        sb.WriteString(escapeLabel(n.Name))
    }
    if !isZero(n.BranchLength) {
        sb.WriteString(":" + formatFloat(n.BranchLength))
    }
}

func formatFloat(f float64) string {
    return strconv.FormatFloat(f, 'f', -1, 64)
}

func isZero(f float64) bool { return f == 0 }

func escapeLabel(s string) string {
    // minimal escaping: wrap in quotes if contains special chars
    for _, r := range s {
        if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' || r == '.') {
            return "'" + strings.ReplaceAll(s, "'", "\\'") + "'"
        }
    }
    return s
}

// newickParser is a tiny recursive descent parser
type newickParser struct {
    input string
    pos   int
}

func (p *newickParser) peek() rune {
    if p.pos >= len(p.input) {
        return 0
    }
    return rune(p.input[p.pos])
}

func (p *newickParser) next() rune {
    if p.pos >= len(p.input) {
        return 0
    }
    r := rune(p.input[p.pos])
    p.pos++
    return r
}

func (p *newickParser) skipWhitespace() {
    for unicode.IsSpace(p.peek()) {
        p.next()
    }
}

func (p *newickParser) parseSubtree() (*Node, error) {
    p.skipWhitespace()
    if p.peek() == '(' {
        // internal node with children
        p.next() // consume '('
        root := NewNode("", 0)
        for {
            child, err := p.parseSubtree()
            if err != nil {
                return nil, err
            }
            root.AddChild(child)
            p.skipWhitespace()
            if p.peek() == ',' {
                p.next()
                continue
            }
            if p.peek() == ')' {
                p.next()
                break
            }
            return nil, fmt.Errorf("expected ',' or ')' at pos %d", p.pos)
        }
        // maybe name and branch length
        p.skipWhitespace()
        label, _ := p.parseLabel()
        if label != "" {
            root.Name = label
        }
        p.skipWhitespace()
        if p.peek() == ':' {
            p.next()
            bl, err := p.parseNumber()
            if err != nil {
                return nil, err
            }
            root.BranchLength = bl
        }
        return root, nil
    }
    // leaf node: label[:branch]
    label, ok := p.parseLabel()
    if !ok {
        return nil, fmt.Errorf("expected label at pos %d", p.pos)
    }
    n := NewNode(label, 0)
    p.skipWhitespace()
    if p.peek() == ':' {
        p.next()
        bl, err := p.parseNumber()
        if err != nil {
            return nil, err
        }
        n.BranchLength = bl
    }
    return n, nil
}

func (p *newickParser) parseLabel() (string, bool) {
    p.skipWhitespace()
    if p.peek() == '\'' {
        // quoted label
        p.next()
        start := p.pos
        for {
            ch := p.next()
            if ch == '\'' || ch == 0 {
                break
            }
        }
        lab := p.input[start : p.pos-1]
        return strings.ReplaceAll(lab, "\\'", "'"), true
    }
    // unquoted label: run of non-special chars
    start := p.pos
    for {
        r := p.peek()
        if r == 0 || r == ':' || r == ',' || r == ')' || r == '(' || unicode.IsSpace(r) {
            break
        }
        p.next()
    }
    if p.pos == start {
        return "", false
    }
    return p.input[start:p.pos], true
}

func (p *newickParser) parseNumber() (float64, error) {
    p.skipWhitespace()
    start := p.pos
    // accept digits, dot, exponent
    for {
        r := p.peek()
        if (r >= '0' && r <= '9') || r == '.' || r == 'e' || r == 'E' || r == '+' || r == '-' {
            p.next()
            continue
        }
        break
    }
    if p.pos == start {
        return 0, fmt.Errorf("expected number at pos %d", p.pos)
    }
    s := p.input[start:p.pos]
    v, err := strconv.ParseFloat(s, 64)
    if err != nil {
        return 0, err
    }
    return v, nil
}
