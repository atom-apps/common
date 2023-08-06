package db

import (
	"database/sql/driver"

	"github.com/samber/lo"
)

func ToValuer[T any](v []T) []driver.Valuer {
	return lo.Map(v, func(item T, index int) driver.Valuer {
		return any(item).(driver.Valuer)
	})
}
