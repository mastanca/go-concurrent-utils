package map_reduce

func Map(list []string, fn func(string) int) []int {
	result := make([]int, len(list))
	for i, s := range list {
		result[i] = fn(s)
	}
	return result
}

func Reduce(list []int, fn func(int, int) int) int {
	var result int
	for _, n := range list {
		result = fn(result, n)
	}
	return result
}
