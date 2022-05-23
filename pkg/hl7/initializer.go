package hl7

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ProcessorConfig struct {
	MsgConfig MessageConfig
}

func (pc *ProcessorConfig) LoadConfig() {
	file, err := os.Open("/home/paulo/Projects/Med/pkg/hl7/testconf.json")
	HandleError(err)
	conf, err := ioutil.ReadAll(file)
	HandleError(err)
	json.Unmarshal(conf, &pc.MsgConfig)
}
