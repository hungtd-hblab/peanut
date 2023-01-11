package array

type ArrayElement interface {
	int | float64 | string
}

func IsIn[T ArrayElement](e T, arr []T) bool {
	for _, v := range arr {
		if e == v {
			return true
		}
	}
	return false
}
