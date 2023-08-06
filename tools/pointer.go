package tools

func ToPtr[T any](v T) *T {
	return &v
}

func FromPtr[T any](p *T, defVal T) T {
	if p != nil {
		return *p
	}

	return defVal
}
