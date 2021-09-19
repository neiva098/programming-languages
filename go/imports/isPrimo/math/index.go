package math

func IsPrimo(number int) bool {
	halfNumber := number / 2

	for i := 2; i <= halfNumber; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number == 1 {
		return false
	} else {
		return true
	}
}
