# README - ConsecutiveSum Utils

This repository contains a Go package named `utils` that provides two functions for calculating consecutive sums - `ConsecutiveSumSync` and `ConsecutiveSumAsync`.

## ConsecutiveSumSync

### Description

`ConsecutiveSumSync` is a synchronous function that calculates the sum of consecutive integers from 1 to the given input number (`num`) inclusive.

### Function Signature

```go
func ConsecutiveSumSync(num int64) int64
```

### Usage

```go
package main

import (
	"fmt"
	"github.com/yourusername/your-repo/utils"
)

func main() {
	num := int64(100)
	sum := utils.ConsecutiveSumSync(num)
	fmt.Printf("Consecutive sum of numbers from 1 to %d is %d\n", num, sum)
}
```

## ConsecutiveSumAsync

### Description

`ConsecutiveSumAsync` is an asynchronous function that calculates the sum of consecutive integers from 1 to the given input number (`num`) inclusive, utilizing goroutines for parallel computation.

### Function Signature

```go
func ConsecutiveSumAsync(num int64) int64
```

### Usage

```go
package main

import (
	"fmt"
	"github.com/yourusername/your-repo/utils"
)

func main() {
	num := int64(100)
	sum := utils.ConsecutiveSumAsync(num)
	fmt.Printf("Consecutive sum of numbers from 1 to %d (async) is %d\n", num, sum)
}
```

### Concurrency

The `ConsecutiveSumAsync` function uses goroutines to divide the range of numbers into multiple parts and calculates the sum concurrently, making it suitable for processing large input values efficiently.

### Goroutine Count

By default, this function uses 100 goroutines for parallel computation. You can adjust the number of goroutines by modifying the `goroutines` variable in the code.

## Contributing

If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear and concise messages.
4. Push your changes to your fork.
5. Create a pull request to the main repository.

## Issues

If you encounter any issues or have suggestions for improvement, please open an issue on this repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to replace placeholders with your actual GitHub username and repository name. Enjoy using the `ConsecutiveSum` utilities in your Go projects!