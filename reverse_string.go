package reverse_string

func ReverseString(input string) (output string) {
	inputSlice := []rune(input)
	strLen := len(input) - 1
	var outputSlice []rune
	for i := strLen; i > -1; i-- {
		outputSlice = append(outputSlice, inputSlice[i])
	}
	output = string(outputSlice)
	return output
}
