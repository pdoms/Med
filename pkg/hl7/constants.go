package hl7

type Delimeters struct {
	Field     []byte
	Segment   []byte
	MshOffset int
}

func NewDelimiters() *Delimeters {
	return &Delimeters{
		Field:     []byte{124},
		Segment:   []byte{10},
		MshOffset: 5,
	}
}

type SegmentIds struct {
	MSH []byte
}

func NewSegmentIds() *SegmentIds {
	return &SegmentIds{
		MSH: []byte{77, 83, 72},
	}
}

func NewMessageTypesCollecion() []string {
	return []string{
		"ACK",
		"ADR",
		"ADT",
		"ARD",
	}
}

/* type MessageSegmentsLengths map[string]SegmentLengths
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
} */
