package json

import "encoding/json"

func JSONParse(data interface{}) string {
	result, _ := json.MarshalIndent(data, "", " ")
	return string(result)
}
