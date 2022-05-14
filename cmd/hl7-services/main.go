package main

import (
	"Med/pkg/hl7"
	"strings"
)

func main() {

	r := strings.NewReader("Test bit 1, test bit 2 | test bit 3 \n test bit 4, test bit^^&& \n")
	scanner := hl7.NewScanner(r)
	scanner.Scan()
}
