package input

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const maxNumber = 100000

func ScanNumbers() []int {
	numbers := make([]int, 0)

	var input string

	for {
		_, err := fmt.Scanln(&input)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: empty input string")
			continue
		}

		number, err := strconv.Atoi(input)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: input should be an integer")
			continue
		}

		if number > maxNumber || number < -maxNumber {
			fmt.Fprintln(os.Stderr, "Error: input number out of range")
			continue
		}

		numbers = append(numbers, number)
	}

	return numbers
}
