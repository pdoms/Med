package hl7config

import (
	"errors"
	"fmt"
)

type SegType string
type MsgType string
type FieldIdx int
type FieldType string
type SubFieldIdx int
type SubFieldType string

type SubFieldNode struct {
	Type SubFieldType
	Idx  SubFieldIdx
}

type FieldNode struct {
	Type     FieldType
	Idx      FieldIdx
	Children map[SubFieldIdx]*SubFieldNode
}

type SegmentNode struct {
	Type     SegType
	Children map[FieldIdx]*FieldNode
}

type MessageNode struct {
	Type     MsgType
	Children map[SegType]*SegmentNode
}

type Hl7PackRoot struct {
	Children map[MsgType]*MessageNode
}

func InitHl7ConfigTree() *Hl7PackRoot {
	var root Hl7PackRoot
	root.Children = make(map[MsgType]*MessageNode)
	return &root
}

func (root *Hl7PackRoot) AddMessage(id MsgType) {
	msg := InitMessageNode()
	msg.Type = id
	msg.Children = make(map[SegType]*SegmentNode)
	root.Children[id] = msg
}

func (root *Hl7PackRoot) AddSegment(msgId MsgType, segId SegType) {
	if _, ok := root.Children[msgId]; ok {
		root.Children[msgId].AddSegment(segId)
	} else {
		root.AddMessage(msgId)
		root.Children[msgId].AddSegment(segId)
	}
}

func (root *Hl7PackRoot) AddField(msgId MsgType, segId SegType, fieldIdx FieldIdx, id FieldType) {
	if msg, ok := root.Children[msgId]; ok {
		if _, ok := msg.Children[segId]; ok {
			msg.Children[segId].AddField(fieldIdx, id)
		} else {
			root.AddSegment(msgId, segId)
			msg.Children[segId].AddField(fieldIdx, id)
		}
	} else {
		root.Children[msgId].AddSegment(segId)
		root.Children[msgId].Children[segId].AddField(fieldIdx, id)
	}
}

func (root *Hl7PackRoot) AddSubfield(msgId MsgType, segId SegType, fieldIdx FieldIdx, fieldType FieldType, subIdx SubFieldIdx, subType SubFieldType) {
	if msg, ok := root.Children[msgId]; ok {
		if seg, ok := msg.Children[segId]; ok {
			if _, ok := seg.Children[fieldIdx]; ok {
				seg.Children[fieldIdx].AddSubfield(subIdx, subType)
			} else {
				root.AddField(msgId, segId, fieldIdx, fieldType)
				seg.Children[fieldIdx].AddSubfield(subIdx, subType)
			}
		} else {
			root.AddField(msgId, segId, fieldIdx, fieldType)
			seg.Children[fieldIdx].AddSubfield(subIdx, subType)
		}
		root.AddField(msgId, segId, fieldIdx, fieldType)
		msg.Children[segId].Children[fieldIdx].AddSubfield(subIdx, subType)
	}

}

func InitMessageNode() *MessageNode {
	var msg MessageNode
	msg.Children = make(map[SegType]*SegmentNode)
	return &msg
}

func (msg *MessageNode) AddSegment(id SegType) {
	seg := InitSegmentNode()
	seg.Type = id
	msg.Children[id] = seg
}

func InitSegmentNode() *SegmentNode {
	var seg SegmentNode
	seg.Children = make(map[FieldIdx]*FieldNode)
	return &seg
}

func (seg *SegmentNode) AddField(idx FieldIdx, id FieldType) {
	field := InitFieldNode()
	field.Idx = idx
	field.Type = id
	seg.Children[idx] = field

}

func InitFieldNode() *FieldNode {
	var field FieldNode
	field.Children = make(map[SubFieldIdx]*SubFieldNode)
	return &field
}

func (field *FieldNode) AddSubfield(subIdx SubFieldIdx, subType SubFieldType) {
	sub := InitSubField()
	sub.Idx = subIdx
	sub.Type = subType
	field.Children[subIdx] = sub
}

func (field *FieldNode) GetSubfield(idx SubFieldIdx) (*SubFieldNode, error) {
	if val, ok := field.Children[idx]; ok {
		return val, nil
	} else {
		return &SubFieldNode{}, notFound("subfield")
	}

}

func (field *FieldNode) GetSubfieldByType(id SubFieldType) *SubFieldNode {
	for item := range field.Children {
		fmt.Println(item)
	}
	return &SubFieldNode{}

}

func InitSubField() *SubFieldNode {
	var sub SubFieldNode
	return &sub
}

func notFound(key string) error {
	return errors.New(fmt.Sprintf("error: %s not found", key))
}
