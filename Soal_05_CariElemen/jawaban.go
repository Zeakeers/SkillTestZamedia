func IndexOf(numbers []int, target int) int {
	for i, v := range numbers {
		if v == target {
			return i
		}
	}
	return -1
}
