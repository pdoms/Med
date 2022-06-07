package hl7

import "testing"

var RAW = "MSH|^~\\&|EPICADT|DH|LABADT|DH||ADT^A01|HL7MSG00001|P|2.3|\n"

func TestUnmarshal(t *testing.T) {
	var msg Message
	r := Unmarshal([]byte(RAW), &msg)
	if r != nil {
		t.Errorf("Unmarshalling failed")
	}
}
