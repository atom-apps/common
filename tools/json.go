package tools

import "encoding/json"

func JsonStr[T any](val T) string {
	if b, err := json.Marshal(val); err != nil {
		return ""
	} else {
		return string(b)
	}
}

func JsonUnmarshal[T any](data []byte) (T, error) {
	var item T
	err := json.Unmarshal(data, &item)
	return item, err
}

func JsonMustUnmarshal[T any](data []byte) T {
	var item T
	err := json.Unmarshal(data, &item)
	if err != nil {
		return item
	}

	return item
}

func JsonMustUnmarshalStr[T any](data string) T {
	return JsonMustUnmarshal[T]([]byte(data))
}
