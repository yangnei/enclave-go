package util

import "encoding/json"

func GetPointer[T any](val T) (res *T) {
	res = &val
	return res
}

func MustMarshalIndent[T any](s T) string {
	res, _ := json.MarshalIndent(s, "", "  ")
	return string(res)
}

func NewTradingPair(base, quote string) string {
	return base + "-" + quote
}
