# ConsecutiveSum in Golang

This is a Go (Golang) repository that provides two functions to calculate the sum of consecutive integers from 1 to a given input number (inclusive). The repository contains the following functions:

1. `ConsecutiveSum`: A synchronous function that calculates the sum of consecutive integers in a single-threaded manner.
2. `ConsecutiveSumConcurrent`: A concurrent function that leverages Goroutines to calculate the sum of consecutive integers concurrently.

## ConsecutiveSum

### Description

This function calculates the sum of consecutive integers sequentially. It takes an `int64` parameter `num`, which represents the upper bound of the consecutive integers to sum. It returns the sum as an `int64`.

### Function Signature

```go
func ConsecutiveSum(num int64) int64
```

### Usage

```go
package main

import (
	"fmt"
	"main/utils"
)

func main() {
	num := int64(100)
	sum := utils.ConsecutiveSum(num)
	fmt.Printf("Consecutive sum of numbers from 1 to %d is %d\n", num, sum)
}
```

## ConsecutiveSumConcurrent

### Description

This function calculates the sum of consecutive integers concurrently using Goroutines. It also takes an `int64` parameter `num` and returns the sum as an `int64`. It divides the calculation into multiple Goroutines (controlled by the `goroutines` constant) to improve performance for large values of `num`.


### Function Signature

```go
func ConsecutiveSumConcurrent(num int64) int64
```

### Usage

```go
package main

import (
	"fmt"
	"main/utils"
)

func main() {
	num := int64(100)
	sum := utils.ConsecutiveSumConcurrent(num)
	fmt.Printf("Consecutive sum of numbers from 1 to %d (concurrent) is %d\n", num, sum)
}
```

### Concurrency

The `ConsecutiveSumConcurrent` function uses goroutines to divide the range of numbers into multiple parts and calculates the sum concurrently, making it suitable for processing large input values efficiently.

### Goroutine Count

By default, this function uses 20 goroutines for parallel computation. You can adjust the number of goroutines by modifying the `goroutines` variable in the code.

## Contributing

If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear and concise messages.
4. Push your changes to your fork.
5. Create a pull request to the main repository.

## License

This project is licensed under the MIT License.

---

Enjoy calculating the sum of consecutive integers efficiently with Go! If you have any questions or need further assistance, please don't hesitate to reach out.






