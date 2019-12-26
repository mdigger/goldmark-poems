package poems

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type poemParser struct{}

var defaultParser parser.BlockParser = new(poemParser)

func (*poemParser) Trigger() []byte {
	return nil
}

func (*poemParser) Open(parent ast.Node, reader text.Reader, pc parser.Context) (ast.Node, parser.State) {
	line, segment := reader.PeekLine()
	pos, padding := util.IndentPosition(line, reader.LineOffset(), 4)
	if pos < 0 || util.IsBlank(line) {
		return nil, parser.NoChildren
	}
	node := NewPoem()
	reader.AdvanceAndSetPadding(pos, padding)
	_, segment = reader.PeekLine()
	node.Lines().Append(segment)
	reader.Advance(segment.Len() - 1)
	return node, parser.NoChildren
}

func (*poemParser) Continue(node ast.Node, reader text.Reader, pc parser.Context) parser.State {
	line, segment := reader.PeekLine()
	if util.IsBlank(line) {
		node.Lines().Append(segment.TrimLeftSpaceWidth(4, reader.Source()))
		return parser.Continue | parser.NoChildren
	}
	pos, padding := util.IndentPosition(line, reader.LineOffset(), 4)
	if pos < 0 {
		return parser.Close
	}
	reader.AdvanceAndSetPadding(pos, padding)
	_, segment = reader.PeekLine()
	node.Lines().Append(segment)
	reader.Advance(segment.Len() - 1)
	return parser.Continue | parser.NoChildren
}

func (*poemParser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	// trim trailing blank lines
	lines := node.Lines()
	length := lines.Len() - 1
	source := reader.Source()
	for length >= 0 {
		line := lines.At(length)
		if util.IsBlank(line.Value(source)) {
			length--
		} else {
			break
		}
	}
	lines.SetSliced(0, length+1)
}

func (*poemParser) CanInterruptParagraph() bool {
	return false
}

func (*poemParser) CanAcceptIndentedLine() bool {
	return true
}
