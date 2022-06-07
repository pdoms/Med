package hl7config

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	fmt.Println("1. Initiate Hl7ConfigTree ------")
	root := InitHl7ConfigTree()
	if root == nil || len(root.Children) != 0 {
		t.Errorf("1. InitConfig Failed")
	}
	fmt.Println("2. Adding Message to root ------")
	root.AddMessage("MSH")
	if root.Children["MSH"] == nil || len(root.Children) == 0 {
		t.Errorf("2. Failed to add message to root")
	}
	fmt.Println("3.1 Adding Segment to root")
	fmt.Println("3.2 Adding Segment to MessageNode")

}
