package hl7

import (
	"bytes"
	"fmt"
	"log"
)

type Hl7MESSAGE struct {
	Config ProcessorConfig
	Raw    []byte
	MSH    []byte
	Type   string
	Delims *Delimeters
}

func (msg *Hl7MESSAGE) UnmarshalHl7(raw []byte) {
	msg.Raw = raw
	msg.Delims = NewDelimiters()
	temp := bytes.Split(msg.Raw, msg.Delims.Segment)
	msg.parseMSH(temp[0])

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
		fmt.Println(msg.Type)
		replace := make([]byte, 4)
		//replace escape characters
		copy(msg.MSH[4:], replace[:])
		fields := bytes.Split(msg.MSH, msg.Delims.Field)
		fmt.Println(fields[0])
	}
}

func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
