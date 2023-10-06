package main

import (
	"fmt"
	"main/utils"
	"time"
)

func main() {

	var num int64 = 1_000_000_000
	var sum int64

	fmt.Println("Calculating the consecutive summation from 1 to 1_000_000_000 concurrently...")
	start := time.Now()
	sum = utils.ConsecutiveSumConcurrent(num)
	duration := time.Since(start)

	fmt.Printf("The sum is %d\n", sum)

	fmt.Printf("Concurrently, it took %fs\n", duration.Seconds())

	fmt.Println("\nCalculating the consecutive summation from 1 to 1_000_000_000 synchronously...")
	start = time.Now()
	utils.ConsecutiveSum(num)
	duration = time.Since(start)

	fmt.Printf("The sum is %d\n", sum)
	fmt.Printf("Synchronously, it took %fs\n", duration.Seconds())
}
