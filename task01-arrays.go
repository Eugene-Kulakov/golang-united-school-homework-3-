package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	for _, item := range input {
		result += item
	}
	result /= 15
	return
}
