package hl7config

import (
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

func (root *Hl7PackRoot) GetMessage(msg MsgType) (*MessageNode, error) {
	if message, ok := root.Children[msg]; ok {
		return message, nil
	}
	return nil, notFound("Message")
}

func (root *Hl7PackRoot) AddSegment(msgId MsgType, segId SegType) {
	if _, ok := root.Children[msgId]; ok {
		root.Children[msgId].AddSegment(segId)
	} else {
		root.AddMessage(msgId)
		root.Children[msgId].AddSegment(segId)
	}
}

func (root *Hl7PackRoot) GetSegment(msgId MsgType, segId SegType) (*SegmentNode, error) {
	msg, err := root.GetMessage(msgId)
	if err != nil {
		return nil, err
	}
	return msg.GetSegment(segId)
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

func (root *Hl7PackRoot) GetField(msgId MsgType, segId SegType, identifier interface{}) (*FieldNode, error) {
	seg, err := root.GetSegment(msgId, segId)
	if err != nil {
		return nil, err
	}
	return seg.GetField(identifier)
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

func (root *Hl7PackRoot) GetSubfield(msgId MsgType, segId SegType, fieldIdentifier interface{}, subIdentifier interface{}) (*SubFieldNode, error) {
	field, err := root.GetField(msgId, segId, fieldIdentifier)
	if err != nil {
		return nil, notFound("SubField")
	}
	return field.GetSubfield(subIdentifier)

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

func (msg *MessageNode) GetSegment(segId SegType) (*SegmentNode, error) {
	if seg, ok := msg.Children[segId]; ok {
		return seg, nil
	} else {
		return nil, notFound("Segment")
	}

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

func (seg *SegmentNode) GetField(identifier interface{}) (*FieldNode, error) {
	t, v := getIndex(identifier)
	if t == "fieldType" || t == "type" {
		for field := range seg.Children {
			if seg.Children[field].Type == v {
				return seg.Children[field], nil
			}
		}

	}
	if t == "fieldIdx" || t == "idx" {
		if field, ok := seg.Children[v.(FieldIdx)]; ok {
			return field, nil
		}
	}
	return nil, notFound("Field")
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

func (field *FieldNode) GetSubfield(identifier interface{}) (*SubFieldNode, error) {
	t, v := getIndex(identifier)
	if t == "subFieldIdx" || t == "idx" {
		if val, ok := field.Children[v.(SubFieldIdx)]; ok {
			return val, nil
		}
	}
	if t == "subFieldType" || t == "type" {
		for sub := range field.Children {
			if field.Children[sub].Type == v {
				return field.Children[sub], nil
			}
		}
	}

	return nil, notFound("subfield")

}

func InitSubField() *SubFieldNode {
	var sub SubFieldNode
	return &sub
}

func notFound(key string) error {
	return fmt.Errorf(fmt.Sprintf("error: %s not found", key))
}

func getIndex(id interface{}) (string, interface{}) {
	switch v := id.(type) {
	case FieldIdx:
		return "fieldIdx", v
	case FieldType:
		return "fieldType", v
	case SubFieldIdx:
		return "subFieldIdx", v
	case SubFieldType:
		return "subFieldType", v
	case string:
		return "type", v
	case int:
		return "idx", v
	}
	return "", nil
}
