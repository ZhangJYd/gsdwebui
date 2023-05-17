package utils

func P[t any](v t) *t {
	return &v
}
