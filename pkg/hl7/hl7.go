package hl7

import (
	"bytes"
	"log"
)

type Hl7MESSAGE struct {
	Config      ProcessorConfig
	Raw         []byte
	MSH         []byte
	Type        string
	Delims      *Delimeters
	segmentsRaw [][]byte
}

type fieldCollect struct {
	fields []*segField
}

func (msg *Hl7MESSAGE) UnmarshalHl7(raw []byte) {
	msg.Raw = raw
	msg.Delims = NewDelimiters()
	msg.segmentsRaw = bytes.Split(msg.Raw, msg.Delims.Segment)
	msg.parseMSH(msg.segmentsRaw[0])

}

func (msg *Hl7MESSAGE) DetermineMessageType() {
	types := NewMessageTypesCollecion()
	for _, msgType := range types {
		if bytes.Contains(msg.MSH, []byte(msgType)) {
			msg.Type = msgType
			return
		}
	}
}

func isMHS(first []byte) bool {
	return bytes.Equal(first[:3], NewSegmentIds().MSH)
}

func (msg *Hl7MESSAGE) parseMSH(first []byte) {
	//watch out: if not first where then?
	if isMHS(first) {
		msg.MSH = first
		msg.DetermineMessageType()
		replace := make([]byte, 4)
		//replace escape characters
		copy(msg.MSH[4:], replace[:])
		iter := newFieldIterator(msg.MSH[9:], "MSH", msg.Type, 2)
		lengths := NewMsgSegmentsLength()
		segLength := lengths[msg.Type]["MSH"] + 1
		var col fieldCollect
		col.fields = make([]*segField, segLength)
		prepareMSH(&col)
		iter.iterate(&col)
	}
}

func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type fieldIter struct {
	fieldId     int
	segment     []byte
	segmentId   string
	messageType string
}

type segField struct {
	content []byte
	idx     int
	seg     string
	id      string
}

func (fi *fieldIter) iterate(collection *fieldCollect) {
	for fi.hasNext() {
		n := fi.next()
		collection.fields[n.idx] = n
	}
	n := &segField{
		content: fi.segment,
		idx:     fi.fieldId + 1,
		seg:     fi.segmentId,
	}
	collection.fields[n.idx] = n
}

func (fi *fieldIter) hasNext() bool {
	return bytes.IndexByte(fi.segment, byte('|')) > -1

}

func (fi *fieldIter) next() *segField {
	if fi.hasNext() {
		idx := bytes.IndexByte(fi.segment, byte('|'))
		fi.fieldId++
		n := &segField{
			content: fi.segment[:idx],
			idx:     fi.fieldId,
			seg:     fi.segmentId,
		}
		fi.segment = fi.segment[idx+1:]
		return n
	} else {
		return nil
	}
}

func newFieldIterator(seg []byte, id, msgtype string, startField int) fieldIter {
	return fieldIter{
		segmentId:   id,
		segment:     seg,
		messageType: msgtype,
		fieldId:     startField,
	}
}

func prepareMSH(coll *fieldCollect) {
	coll.fields[1] = &segField{
		content: make([]byte, 1),
		idx:     1,
		seg:     "MSH",
		id:      "MSH.1",
	}
	coll.fields[2] = &segField{
		content: make([]byte, 3),
		idx:     2,
		seg:     "MSH",
		id:      "MSH.1",
	}
}
