package bmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

func TestPlainUsername(t *testing.T) {
	uut := newMatrixUsername("MyUser")

	assert.Equal(t, "MyUser", uut.formatted)
	assert.Equal(t, "MyUser", uut.plain)
}

func TestHTMLUsername(t *testing.T) {
	uut := newMatrixUsername("<b>MyUser</b>")

	assert.Equal(t, "<b>MyUser</b>", uut.formatted)
	assert.Equal(t, "MyUser", uut.plain)
}

func TestFancyUsername(t *testing.T) {
	uut := newMatrixUsername("<MyUser>")

	assert.Equal(t, "&lt;MyUser&gt;", uut.formatted)
	assert.Equal(t, "<MyUser>", uut.plain)
}

func TestMatrixParentIDForReply(t *testing.T) {
	parentID := matrixParentIDForRelation(&event.RelatesTo{
		InReplyTo: &event.InReplyTo{EventID: id.EventID("$reply")},
	})

	assert.Equal(t, "$reply", parentID)
}

func TestMatrixParentIDForThread(t *testing.T) {
	parentID := matrixParentIDForRelation(&event.RelatesTo{
		Type:    event.RelThread,
		EventID: id.EventID("$thread-root"),
		InReplyTo: &event.InReplyTo{
			EventID: id.EventID("$fallback-reply"),
		},
		IsFallingBack: true,
	})

	assert.Equal(t, "$thread-root", parentID)
}
