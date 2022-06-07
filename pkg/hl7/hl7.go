package hl7

import (
	"Med/pkg/hl7config"
	"bytes"
)

type Msg string
type Seg int
type Fld int
type Sub int

type Message struct {
	segmentsRaw [][]byte
	Type        Msg
	Segments    map[Seg]*Segment
	config      *hl7config.Hl7PackRoot
}

type Segment struct {
	Type   string
	Fields map[Fld]*Field
}

type Field struct {
	Type      string
	Idx       string
	Value     string
	SubFields map[Sub]*SubField
}

type SubField struct {
	Type  string
	Idx   string
	Value string
}

func Unmarshal(msg []byte, target *Message) error {
	delims := NewDelimiters()
	target.segmentsRaw = bytes.Split(msg, delims.Segment)
	target.config = hl7config.Load()
	for _, seg := range target.segmentsRaw {
		segCopy := make([]byte, len(seg))
		copy(segCopy[:], seg)
		iterator := NewSegmentIterator(segCopy)
		iterator.Iterate(target)

	}

	return nil
}
