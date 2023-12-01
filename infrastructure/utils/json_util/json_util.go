package json_util

import "encoding/json"

// StructToJSON 将结构体转换为 JSON 字符串
func StructToJSON(data interface{}) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	jsonString := string(jsonBytes)
	return jsonString, nil
}

func NoErrToJSON(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	jsonString := string(jsonBytes)
	return jsonString
}
