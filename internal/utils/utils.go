package utils

import "log"

func SliceContainsString(s []string, item string) bool {
	for _, a := range s {
		if a == item {
			return true
		}
	}
	return false
}

func HandleError(err error) {
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}
}
