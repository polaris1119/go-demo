package mention

import (
	gast "github.com/yuin/goldmark/ast"
)

// KindMention is a NodeKind of the Mention node.
var KindMention = gast.NewNodeKind("Mention")

type MentionNode struct {
	gast.BaseInline
	Who string
}

// NewStrikethrough returns a new Mention node.
func NewMentionNode(username string) *MentionNode {
	return &MentionNode{
		BaseInline: gast.BaseInline{},
		Who:        username,
	}
}

// Dump implements Node.Dump.
func (n *MentionNode) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// Kind implements Node.Kind.
func (n *MentionNode) Kind() gast.NodeKind {
	return KindMention
}
