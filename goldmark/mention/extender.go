package mention

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type mention struct{}

var Mention = mention{}

func (m mention) Extend(markdown goldmark.Markdown) {
	markdown.Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(NewMentionParser(), 500),
	))

	markdown.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewMentionHTMLRenderer(), 500),
	))
}
