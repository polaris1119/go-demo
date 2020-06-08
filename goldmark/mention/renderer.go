package mention

import (
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type mentionHTMLRenderer struct{}

func NewMentionHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	return mentionHTMLRenderer{}
}

func (m mentionHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindMention, m.renderMention)
}

func (m mentionHTMLRenderer) renderMention(w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering {
		mn := n.(*MentionNode)
		w.WriteString(`<a href="https://studygolang.com/user/` + mn.Who + `">@`)
		w.WriteString(mn.Who)
	} else {
		w.WriteString("</a>")
	}

	return gast.WalkContinue, nil
}
