package array_slices

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, v := range numbers {
		sum = sum + v
	}
	return
}
