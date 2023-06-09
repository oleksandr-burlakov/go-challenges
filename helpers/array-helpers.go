package helpers

type Number interface {
	int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64
}

func Filter[T any](arr []T, test func(T) bool) (ret []T) {
	for _, e := range arr {
		if test(e) {
			ret = append(ret, e)
		}
	}
	return ret
}

func Max[T Number](array []T) (res T) {
	for i, item := range array {
		if i == 0 || item > res {
			res = item
		}
	}
	return res
}

func Min[T Number](array []T) (res T) {
	for i, item := range array {
		if i == 0 || item < res {
			res = item
		}
	}
	return res
}
