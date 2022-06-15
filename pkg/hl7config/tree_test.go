package hl7config

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	fmt.Println("Test 1: Initiate Hl7ConfigTree ---------------------------------------------")
	root := InitHl7ConfigTree()
	if root == nil || len(root.Children) != 0 {
		t.Errorf("1. InitConfig Failed")
	}
	fmt.Println("Test 2: Adding Message to root ---------------------------------------------")
	root.AddMessage("ADT")
	if root.Children["ADT"] == nil || len(root.Children) == 0 {
		t.Errorf("2. Failed to add message to root")
	}
	fmt.Println("Test 3: Getting message via root -------------------------------------------")
	msg, err := root.GetMessage("ADT")
	if err != nil {
		t.Errorf("3-1. Failed to get message via root")
	}
	if msg.Type != "ADT" {
		t.Errorf("3-1. Wanted 'ADT' message got %v", msg.Type)
	}
	fmt.Println("Test 4: Error (no found) getting message via root---------------------------")
	msg_err, err := root.GetMessage("ATD")
	if err == nil || msg_err != nil {
		t.Errorf("4. Wanted 'error: Message not found' got %v for error handle", err)
	} else {
		fmt.Println("  ---> Received expected error message of: ", err)
	}

	fmt.Println("Test 5: Add segment via root -----------------------------------------------")
	message := MsgType("ORM")
	segment := SegType("PID")
	root.AddSegment(message, segment)
	res_1, err := root.GetMessage("ORM")
	if err != nil {
		t.Errorf("5-1. Wanted %v got %v", message, res_1)
	}
	resMsgType := res_1.Type
	_, ok := res_1.Children[segment]
	if resMsgType != message {
		t.Errorf("5-1. Wanted %v got %v", message, resMsgType)
	}
	if !ok {
		t.Errorf("5-2. Wanted true got %v", ok)
	}

	fmt.Println("Test 6: Add segment via MessageNode ----------------------------------------")
	res_1.AddSegment("MSH")
	if len(res_1.Children) != 2 {
		t.Errorf("6. Wanted 2 items in children, got %v", ok)

	}
	fmt.Println("Test 7: Get segment via root -----------------------------------------------")
	msh_seg, err := root.GetSegment("ORM", "MSH")
	if msh_seg == nil || err != nil || msh_seg.Type != "MSH" {
		t.Errorf("7. Wanted Segment got %v", err)
	}
	fmt.Println("Test 8: Get segment via MessageNode ----------------------------------------")
	msg_seg_from_msg, err := res_1.GetSegment("MSH")
	if msg_seg_from_msg == nil || err != nil || msg_seg_from_msg.Type != "MSH" {
		t.Errorf("8. Wanted Segment got %v", err)
	}
	fmt.Println("Test 9: Random Error check -------------------------------------------------")
	msh_seg_err, err := root.GetSegment("OM", "MSH")
	if msh_seg_err != nil || err == nil {
		t.Errorf("9. Wanted 'error: Message not found' got %v for error handle", err)
	} else {
		fmt.Println("  ---> Received expected error message of: ", err)
	}

	var msg_t MsgType = "ADT"
	var seg SegType = "PID"
	var field FieldType = "PatientId"
	var fieldIdx FieldIdx = 2
	var fieldType2 FieldType = "PatientName"
	var fieldIdx2 FieldIdx = 3

	fmt.Println("Test 10. Add Field via root ------------------------------------------------")
	root.AddField(msg_t, seg, fieldIdx, field)
	pid, err := root.GetSegment(msg_t, seg)
	if err != nil || pid.Type != seg {
		t.Errorf("10. Wanted %v got %v", seg, pid.Type)
	}
	fmt.Println("Test 11. Get Field via root by Identifier ----------------------------------")
	f, err := root.GetField(msg_t, seg, fieldIdx)
	if err != nil {
		t.Errorf("11-1. Got error getting field by idx")
	}
	f2, err := root.GetField(msg_t, seg, field)
	if err != nil {
		t.Errorf("11-2. Got error getting field by Type")
	}
	if f.Type != f2.Type {
		t.Errorf("11-3 Field names do not match")
	}

	fmt.Println("Test 12. Add Field via Segment Node ----------------------------------------")
	pid.AddField(fieldIdx2, fieldType2)
	fmt.Println("Test 13. Get Field via Segment Node by Identifier --------------------------")
	f3, err := pid.GetField(fieldIdx2)
	if err != nil {
		t.Errorf("13-1. Got error getting field by idx")
	}
	f4, err := pid.GetField(fieldType2)
	if err != nil {
		t.Errorf("13-2. Got error getting field by Type")
	}
	if f3.Type != f4.Type {
		t.Errorf("13-3 Field names do not match")
	}

	var subIdx1 SubFieldIdx = 1
	var subT1 SubFieldType = "FirstName"
	var subIdx2 SubFieldIdx = 2
	var subT2 SubFieldType = "LastName"

	fmt.Println("Test 14. Add Subfield via root ---------------------------------------------")
	root.AddSubfield(msg_t, seg, fieldIdx2, fieldType2, subIdx1, subT1)
	retrField, err := root.GetField(msg_t, seg, fieldIdx2)
	if err != nil || retrField.Type != fieldType2 {
		t.Errorf("14. Wanted %v got %v", seg, pid.Type)
	}
	fmt.Println("Test 15. Get subfield via root by Identifier -------------------------------")
	retrSub1, err := retrField.GetSubfield(subIdx1)
	if err != nil {
		t.Errorf("15-1. Got error getting sub field by idx")
	}
	retrSub2, err := retrField.GetSubfield(subT1)
	if err != nil {
		t.Errorf("15-2. Got error getting sub field by Type")
	}
	if retrSub1.Type != retrSub2.Type {
		t.Errorf("15-3. Sub Field types do not match.")
	}
	fmt.Println("Test 16. Add Subfield via fieldNode ----------------------------------------")
	retrField.AddSubfield(subIdx2, subT2)
	fmt.Println("Test 17. Get subfield via fieldNode by Identifier --------------------------")
	retrSub3, err := retrField.GetSubfield(subIdx2)
	if err != nil {
		t.Errorf("17-1. Got error getting sub field by idx")
	}
	retrSub4, err := retrField.GetSubfield(subT2)
	if err != nil {
		t.Errorf("17-2. Got error getting sub field by Type")
	}

	if retrSub3.Type != retrSub4.Type {
		t.Errorf("17-3 Sub field names do not match")
	}
}

func TestGetIdx(t *testing.T) {
	fmt.Println("Test 18. Get Index ---------------------------------------------------------")
	var fieldIdx FieldIdx = 1
	var fieldIdxInt = 1
	var fieldType FieldType = "fieldType"
	var fieldTypeStr = "fieldType"
	var subfieldIdx SubFieldIdx = 1
	var SubFieldType SubFieldType = "subfieldType"
	t1, v1 := getIndex(fieldIdx)
	if t1 != "fieldIdx" && v1 != 1 {
		t.Errorf("18-1-1. Wanted %v got %v (type)", "fieldIdx", t1)
		t.Errorf("18-1-2. Wanted %v got %v (value)", 1, v1)
	}
	t2, v2 := getIndex(fieldIdxInt)
	if t2 != "idx" && v2 != 1 {
		t.Errorf("18-2-1. Wanted %v got %v (type)", "idx", t2)
		t.Errorf("18-2-2. Wanted %v got %v (value)", 1, v1)
	}
	t3, v3 := getIndex(fieldType)
	if t3 != "fieldType" && v3 != "fieldType" {
		t.Errorf("18-3-1. Wanted %v got %v (type)", "fieldType", t3)
		t.Errorf("18-3-2. Wanted %v got %v (value)", "fieldType", v3)
	}
	t4, v4 := getIndex(fieldTypeStr)
	if t4 != "type" && v4 != "fieldType" {
		t.Errorf("18-4-1. Wanted %v got %v (type)", "type", t4)
		t.Errorf("18-4-2. Wanted %v got %v (value)", "fieldType", v4)
	}
	t5, v5 := getIndex(subfieldIdx)
	if t5 != "subFieldIdx" && v5 != 1 {
		t.Errorf("18-5-1. Wanted %v got %v (type)", "subFieldType", t5)
		t.Errorf("18-5-2. Wanted %v got %v (value)", 1, v5)
	}
	t6, v6 := getIndex(SubFieldType)
	if t6 != "subFieldType" && v6 != "subfieldType" {
		t.Errorf("18-3-1. Wanted %v got %v (type)", "subFieldType", t6)
		t.Errorf("18-3-2. Wanted %v got %v (value)", "subfieldType", v6)
	}
}
