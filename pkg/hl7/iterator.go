package hl7

import (
	"bytes"
	"fmt"
)

//new iterator will need to have the config

type SegmentIterator struct {
	rawSeg          []byte
	currentFieldIdx int
	currentSegment  string
	Segment         *Segment
}

type iteratorField struct {
	content []byte
	idx     int
}

func (si *SegmentIterator) hasNext() bool {
	return bytes.IndexByte(si.rawSeg, byte('|')) > -1
}

func (si *SegmentIterator) next() *iteratorField {
	if si.hasNext() {
		idx := bytes.IndexByte(si.rawSeg, byte('|'))
		si.currentFieldIdx++
		field := &iteratorField{
			content: si.rawSeg[:idx],
			idx:     si.currentFieldIdx,
		}
		si.rawSeg = si.rawSeg[idx+1:]
		return field
	} else {
		return nil
	}
}

func NewSegmentIterator(rawSeg []byte) *SegmentIterator {
	if len(rawSeg) < 3 {
		return &SegmentIterator{}
	}
	fmt.Println("RAW", rawSeg)
	segType := string(rawSeg[:3])
	return &SegmentIterator{
		currentSegment:  segType,
		rawSeg:          rawSeg,
		currentFieldIdx: 0,
		Segment: &Segment{
			Type:   segType,
			Fields: make(map[Fld]*Field),
		},
	}
}

func (si *SegmentIterator) Iterate(msg *Message) {
	for si.hasNext() {
		fmt.Println(si.currentSegment)
		next := si.next()

		fmt.Println("NEXT", next)

	}
}
