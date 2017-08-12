package gray

//GenerateGrayCode will create all the gray code values for n-bit size recursivly.
func GenerateGrayCode(size int) (grayCode []string) {
	if size > 1 {
		grayCode = GenerateGrayCode(size - 1)
		grayCode = append(grayCode, reverse(grayCode)...)
		for i, j := 0, len(grayCode)/2; i < len(grayCode)/2; i, j = i+1, j+1 {
			grayCode[i] = "0" + grayCode[i]
			grayCode[j] = "1" + grayCode[j]
		}
	} else {
		grayCode = []string{"0", "1"}
	}
	return
}

func reverse(numbers []string) []string {
	newNumbers := make([]string, len(numbers))
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}
	return newNumbers
}
