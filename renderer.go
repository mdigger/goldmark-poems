package poems

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type render struct {
	start, end string
}

var defaulRenderer = &render{
	start: "<div class=\"poem\">\n",
	end:   "</div>\n",
}

func (r *render) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindPoem, r.renderPoem)
}

func (r *render) renderPoem(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		w.WriteString(r.start)
	} else {
		w.WriteString(r.end)
	}
	return ast.WalkContinue, nil
}
