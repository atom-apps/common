package db

import (
	"fmt"
)

func Like(raw string) string {
	return fmt.Sprintf("%%%s%%", raw)
}

func LLike(raw string) string {
	return fmt.Sprintf("%%%s", raw)
}

func RLike(raw string) string {
	return fmt.Sprintf("%s%%", raw)
}
