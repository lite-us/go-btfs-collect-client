package logclient

import (
	"github.com/steveyeom/go-btfs-logclient/logproto"
	"fmt"
)

type Entry interface {
	Type() int
	Value() interface{}
	SetValue(interface{}) error
}

const (
	LINE_ENTRY = iota + 1
	PROTO_ENTRY
)

type LineEntry struct {
	Text string
}

func (line LineEntry) Type() int {
	return LINE_ENTRY
}

func (line LineEntry) Value() interface{} {
	return line.Text
}

func (line LineEntry) SetValue(v interface{}) error {
	ent, ok := v.(LineEntry)
	if !ok {
		return fmt.Errorf("expected LineEntry type, but got %T", v)
	}
	line.Text = ent.Text
	return nil
}

type ProtoEntry struct {
	Pentry *logproto.Entry
}

func (pe ProtoEntry) Type() int {
	return PROTO_ENTRY
}

func (pe ProtoEntry) Value() interface{} {
	return pe.Pentry
}

func (pe ProtoEntry) SetValue(v interface{}) error {
	ent, ok := v.(*logproto.Entry)
	if !ok {
		return fmt.Errorf("expected *logproto.Entry type, but got %T", v)
	}
	pe.Pentry = ent
	return nil
}
