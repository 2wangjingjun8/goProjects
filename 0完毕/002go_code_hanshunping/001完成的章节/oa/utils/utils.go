package utils

func Cal(a float64,b float64,operate byte ) float64 {
	var res float64
	switch operate {
		case '+':
			res = a + b
		case '-':
			res = a - b
		case '*':
			res = a * b
		case '/':
			res = a / b
	}
	return res
}