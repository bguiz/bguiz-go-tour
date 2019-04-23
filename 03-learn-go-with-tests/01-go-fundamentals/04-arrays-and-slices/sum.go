package arrays_and_slices

func Sum5Array(numbers [5]int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// NOTE
// - arrays in go are always of fixed size, and hence we don't use them very often
// - `range` returns something similar to a list comprehension
// - in each iteration, the index and the value are returned by `range`
// - in this case, we only care about the value,
//   so we use `_` to indicate that the index is extraneous - called the blank identifier

func SumSlice(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// NOTE
// - unlike arrays, slices can have dynamic size, and are therefore used more often than arrays

func SumAll(numbersToSum ...[]int) (sums []int) {
	sums = make([]int, len(numbersToSum))
	for index, numbers := range numbersToSum {
		sums[index] = SumSlice(numbers)
	}
	return
}

// NOTE
// - `...` in a function parameter list denotes that this function is variadic,
//   which means that it can take in a variable number of arguments
// - the the function specifies not just the return type, but also its name: `(sums []int)`
// - `make` is used to create a slice with a specified starting capacity
// - we can reference an item on a slice by index, for both reading from and assignment to
//   - this only works when that index already exists within the current capacity that the slice has

func SumAllUsingAppend(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlice(numbers))
	}
	return
}

// NOTE
// - use `append` to concatenate two slices
//   - this is necessary when the size of the slice is not known ahead of time

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, SumSlice(numbers[1:]))
		}
	}
	return
}

// NOTE
// - when you reference a slice with a single number, e.g. `aSlice[5]`,
//   it evaluates to a single item
// - when you reference a slice with a range, e.g. `aSlice[1:3]`,
//   it evaluates to a new slice that is a subset of the original slice
//   - when the start/end of the range is not specified it defaults to
//     `0` and `len(aSlice)` respectively
// - slice internals:
//   - https://blog.golang.org/go-slices-usage-and-internals
