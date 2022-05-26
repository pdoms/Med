package hl7

func NewMessageTypesCollecion() []string {
	return []string{
		"ACK",
		"ADR",
		"ADT",
		"ARD",
	}
}

type MessageSegmentsLengths map[string]SegmentLengths
type SegmentLengths map[string]int

func NewMsgSegmentsLength() MessageSegmentsLengths {
	return MessageSegmentsLengths{
		"ACK": SegmentLengths{
			"MSH": 0,
		},
		"ADT": SegmentLengths{
			"MSH": 21,
		},
	}
}
