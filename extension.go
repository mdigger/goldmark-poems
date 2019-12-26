package poems

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// A poemExtension is goldmark extension for poems code block extension in
// markdown.
type poemExtension struct{}

// Extend implement goldmark.Extender interface.
func (*poemExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(defaulTransformer, 100)),
		parser.WithBlockParsers(
			util.Prioritized(defaultParser, 450)),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(defaulRenderer, 100)),
	)
}

// Extension is a initialized goldmark extension for poems code block support.
var Extension goldmark.Extender = new(poemExtension)

// Enable is goldmark.Enable for poem code blocks extension.
var Enable = goldmark.WithExtensions(Extension)
