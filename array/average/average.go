package average

func Calc(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	result := 0.0
	for _, v := range numbers {
		result += v
	}
	result /= float64(len(numbers))
	return result
}
