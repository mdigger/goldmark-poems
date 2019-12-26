package poems

import (
	"github.com/yuin/goldmark/ast"
)

// Poem struct represents a poem block of Markdown text.
type Poem struct {
	ast.BaseBlock
}

// Dump implements Node.Dump .
func (n *Poem) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, nil, nil)
}

// KindPoem is a NodeKind of the Poem node.
var KindPoem = ast.NewNodeKind("Poem")

// Kind implements Node.Kind.
func (n *Poem) Kind() ast.NodeKind {
	return KindPoem
}

// NewPoem returns a new Poem node.
func NewPoem() *Poem {
	return &Poem{
		BaseBlock: ast.BaseBlock{},
	}
}
