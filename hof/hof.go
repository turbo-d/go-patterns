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

func Reduce(slice []int, fn func(accumulator, x int) int) int {
	acc := slice[0]
	for i := 1; i < len(slice); i++ {
		acc = fn(acc, slice[i])
	}
	return acc
}

func Every(slice []int, fn func(int) bool) bool {
	for _, s := range slice {
		if !fn(s) {
			return false
		}
	}
	return true
}

func Some(slice []int, fn func(int) bool) bool {
	for _, s := range slice {
		if fn(s) {
			return true
		}
	}
	return false
}

func Find(slice []int, fn func(int) bool) (int, bool) {
	for _, s := range slice {
		if fn(s) {
			return s, true
		}
	}
	var notFound int
	return notFound, false
}

func FindIndex(slice []int, fn func(int) bool) (int, bool) {
	for i, s := range slice {
		if fn(s) {
			return i, true
		}
	}
	var notFound int
	return notFound, false
}
