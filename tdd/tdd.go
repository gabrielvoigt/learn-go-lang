package tdd

func Multiples(m, n int) []int{
	result := []int{}
	for i := m; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = append(result, i)
		}
	}

	return result
}