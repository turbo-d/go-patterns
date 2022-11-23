package hof

func Map(slice []int, fn func(int) int) []int {
	out := make([]int, len(slice))
	for i, s := range slice {
		out[i] = fn(s)
	}
	return out
}

func Filter(slice []int, fn func(int) bool) []int {
	out := make([]int, 0, len(slice))
	for _, s := range slice {
		if fn(s) {
			out = append(out, s)
		}
	}
	return out
}

func Reduce(slice []int, fn func(a, b int) int) int {
	temp := make([]int, len(slice))
	copy(temp, slice)

	for l := 1; l < len(temp); l = l << 1 {
		stride := l << 1
		for i := 0; i < len(temp)-l; i += stride {
			temp[i] = fn(temp[i], temp[i+l])
		}
	}

	if len(temp)%2 != 0 {
		temp[0] += temp[len(temp)-1]
	}

	return temp[0]
}

func Every(slice []int, fn func(int) bool) bool {
	for _, s := range slice {
		if !fn(s) {
			return false
		}
	}
	return true
}
