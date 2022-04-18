package middleware

import (
	"echo-app/util"
	"regexp"
)

func Validate(structData interface{}) bool {
	mapData := util.StructToMap(structData)
	for key, value := range mapData {
		switch key {
		case "name":
			if m, _ := regexp.MatchString(".{1,50}", value.(string)); !m {
				return false
			}
		case "email":
			if m, _ := regexp.MatchString("[a-zA-Z0-9_]{8,50}@[a-zA-Z0-9._]{8,50}", value.(string)); !m {
				return false
			}
		case "password":
			if m, _ := regexp.MatchString("[a-zA-Z0-9_]{8,50}", value.(string)); !m {
				return false
			}
		default:
		}
	}
	return true
}
