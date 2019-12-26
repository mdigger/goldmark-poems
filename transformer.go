package poems

import (
	"bytes"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type transformer struct{}

var defaulTransformer = new(transformer)

// nbsp sets the indentation element.
var nbsp = []byte("&nbsp;") // &emsp;

func (*transformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	source := reader.Source()
	ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering || node.Kind() != KindPoem {
			return ast.WalkContinue, nil
		}
		for child := node.FirstChild(); child != nil; child = child.NextSibling() {
			if child.Kind() != ast.KindText {
				continue
			}
			textNode := child.(*ast.Text)
			if textNode.SoftLineBreak() && textNode.NextSibling() != nil {
				textNode.SetHardLineBreak(true)
			}
			text := textNode.Text(source)
			if spaces := util.TrimLeftSpaceLength(text); spaces > 0 {
				indent := ast.NewString(bytes.Repeat(nbsp, spaces))
				indent.SetCode(true)
				node.InsertBefore(node, textNode, indent)
				textNode.Segment.Start += spaces - textNode.Segment.Padding
				textNode.Segment.Padding = 0
			}
		}
		return ast.WalkSkipChildren, nil
	})
}
