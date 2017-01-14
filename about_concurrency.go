package go_koans

func isPrimeNumber(possiblePrime int) bool {
	if possiblePrime == 2 {
		return true
	} else if possiblePrime%2 == 0 {
		return false
	} else {
		for underPrime := 3; underPrime < 1+possiblePrime/2; underPrime += 2 {
			if possiblePrime%underPrime == 0 {
				return false
			}
		}
		return true
	}
}

func findPrimeNumbers(channel chan int) {
	for i := 2; ; /* infinite loop */ i++ {
		// your code goes here
		if isPrimeNumber(i) {
			channel <- i
		}
		assert(i < 200) // i is afraid of heights
	}
}

func aboutConcurrency() {
	ch := make(chan int)

	go findPrimeNumbers(ch)

	assert(<-ch == 2)
	assert(<-ch == 3)
	assert(<-ch == 5)
	assert(<-ch == 7)
	assert(<-ch == 11)
	assert(<-ch == 13)
	assert(<-ch == 17)
	assert(<-ch == 19)
	assert(<-ch == 23)
	assert(<-ch == 29)
	assert(<-ch == 31)
	assert(<-ch == 37)
	assert(<-ch == 41)
	assert(<-ch == 43)
	assert(<-ch == 47)
	assert(<-ch == 53)
	assert(<-ch == 59)
	assert(<-ch == 61)
	assert(<-ch == 67)
	assert(<-ch == 71)
	assert(<-ch == 73)
	assert(<-ch == 79)
	assert(<-ch == 83)
	assert(<-ch == 89)
	assert(<-ch == 97)
}
