package mention

import (
	"regexp"
	"unicode"

	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

var usernameRegexp = regexp.MustCompile(`@([^\s@]{4,20})`)

type mentionParser struct{}

func NewMentionParser() parser.InlineParser {
	return mentionParser{}
}

func (m mentionParser) Trigger() []byte {
	return []byte{'@'}
}

func (m mentionParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	before := block.PrecendingCharacter()
	if !unicode.IsSpace(before) {
		return nil
	}
	line, _ := block.PeekLine()
	matched := usernameRegexp.FindSubmatch(line)
	if len(matched) < 2 {
		return nil
	}
	block.Advance(len(matched[0]))
	node := NewMentionNode(string(matched[1]))
	return node
}
