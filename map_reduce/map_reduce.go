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

func MapStream(list []string, fn func(string) int) <-chan int {
	result := make(chan int, len(list))
	go func() {
		defer close(result)
		for _, s := range list {
			result <- fn(s)
		}
	}()

	return result
}

func ReduceStream(ints <-chan int, fn func(int, int) int) int {
	var result int
	for n := range ints {
		result = fn(result, n)
	}
	return result
}
