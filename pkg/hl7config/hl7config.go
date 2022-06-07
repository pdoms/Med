package hl7config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type CONF map[string]map[string]map[string]interface{}

// type Hl7Entity struct {
// 	Id    string
// 	Index int
// 	Sub   []*Hl7sub
// }

// type Hl7sub struct {
// 	ParentId string
// 	Id       string
// 	Index    int
// }

// type MsgType string
// type SegmentType string

// type Segments map[SegmentType]Hl7Entity

// type MsgConfig map[MsgType]Segments

func Load() *Hl7PackRoot {
	file, err := os.Open("/home/paulo/Projects/Med/pkg/hl7config/test.config.json")
	if err != nil {
		log.Fatalln(err)
	}
	conf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	var config CONF
	json.Unmarshal(conf, &config)
	c := InitHl7ConfigTree()
	for co := range config {
		msgId := MsgType(co)
		c.AddMessage(msgId)
		for seg := range config[co] {
			segId := SegType(seg)
			c.AddSegment(msgId, segId)
			for field := range config[co][seg] {
				pre := config[co][seg][field]
				fieldId := FieldType(field)
				if reflect.ValueOf(pre).Kind() == reflect.String {
					idx := fieldStrToIdx(pre.(string))
					c.AddField(msgId, segId, idx, fieldId)
				} else {
					subfields := pre.(map[string]interface{})
					for sub := range subfields {
						fieldIdx, subFieldIdx := subFieldStrToIdxs(subfields[sub].(string))
						c.AddSubfield(msgId, segId, fieldIdx, fieldId, subFieldIdx, SubFieldType(sub))

					}
				}

			}
		}
	}
	return c
}

func fieldStrToIdx(slotId string) FieldIdx {
	_, after, _ := strings.Cut(string(slotId), ".")
	i, err := strconv.Atoi(after)
	if err != nil {
		log.Fatalln(err)
	}
	return FieldIdx(i)
}

func subFieldStrToIdxs(slotId string) (FieldIdx, SubFieldIdx) {
	split := strings.Split(slotId, ".")
	fIdx, err := strconv.Atoi(split[1])
	if err != nil {
		return FieldIdx(11), SubFieldIdx(-1)
	}
	sIdx, err := strconv.Atoi(split[2])
	if err != nil {
		return FieldIdx(-1), SubFieldIdx(-1)
	}
	return FieldIdx(fIdx), SubFieldIdx(sIdx)
}
