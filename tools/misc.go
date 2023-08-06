package tools

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func BoolToInt[T constraints.Integer](b bool) T {
	if b {
		return T(1)
	}
	return T(0)
}

func StringToInt[T constraints.Integer](str string) T {
	i, err := strconv.Atoi(str)
	if err != nil {
		return T(0)
	}
	return T(i)
}

func IntToString[T constraints.Integer](i T) string {
	return fmt.Sprintf("%d", i)
}

func SliceToString[T any](s []T) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func StringToSlice[T any](s string) ([]T, error) {
	var items []T
	err := json.Unmarshal([]byte(s), &items)
	return items, err
}

func StringToLabels(s string) map[string]string {
	labels := make(map[string]string)
	s = strings.ReplaceAll(s, ",", ";")
	for _, label := range strings.Split(s, ";") {
		kv := strings.Split(label, "=")
		if len(kv) != 2 {
			continue
		}
		labels[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}
	return labels
}

func MapMustString(raw map[string]interface{}) map[string]string {
	if raw == nil {
		return nil
	}

	return lo.MapValues(raw, func(v interface{}, k string) string {
		return fmt.Sprintf("%s", v)
	})
}
