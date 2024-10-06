package util

import (
	"io"
	"os"
	"strconv"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
)

func ParseUint(s string) (uint, error) {
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(u), nil
}

func InverseMap[K, V comparable](d map[K]V) (map[V]K, error) {
	q := make(map[V]K, len(d))
	for k, v := range d {
		if x, ok := q[v]; ok {
			return nil, errorx.Conversion("InverseMap got duplicated key: %v -> %v", x, v)
		}
		q[v] = k
	}
	return q, nil
}

func MustInverseMap[K, V comparable](d map[K]V) map[V]K {
	x, err := InverseMap(d)
	logx.PanicOnError(err)
	return x
}

func ReadAndParse[T any](r io.ReadCloser, f func([]byte) (T, error)) (T, error) {
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		var t T
		return t, err
	}
	return f(b)
}

func OpenAndParse[T any](file string, f func([]byte) (T, error)) (T, error) {
	fp, err := os.Open(file)
	if err != nil {
		var t T
		return t, err
	}
	return ReadAndParse(fp, f)
}
