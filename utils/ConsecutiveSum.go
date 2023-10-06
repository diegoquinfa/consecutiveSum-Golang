package utils

// Is a synchronous function that calculates the sum of consecutive integers from 1 to the given input number (num) inclusive.
func ConsecutiveSum(num int64) int64 {
	if num < 1 {
		return 0
	}

	var sum int64
	for i := int64(0); i <= num; i++ {
		sum += i
	}
	return sum
}

func ConsecutiveSumConcurrent(num int64) int64 {
	if num < 1 {
		return 0
	}

	const goroutines int64 = 20 // Numero de Goroutines
	mod := num % goroutines

	channel := make(chan int64)
	defer close(channel)

	for i := int64(0); i < goroutines; i++ {
		go func(ch chan int64, i int64) {
			var subSum int64
			start := (num/goroutines)*i + 1
			end := (num / goroutines) * (i + 1)

			if i == goroutines-1 {
				end += mod
			}

			for j := start; j <= end; j++ {
				subSum += j
			}
			ch <- subSum
		}(channel, i)
	}

	var sum int64
	for i := int64(0); i < goroutines; i++ {
		sum += <-channel
	}

	return sum
}
