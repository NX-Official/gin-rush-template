package tools

import "encoding/json"

func MustUnmarshal[T any](s string) T {
	var v T
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		panic(err)
	}
	return v
}
