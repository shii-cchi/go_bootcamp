package runner

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
)

func Run(command []string) error {
	var inputData []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			continue
		}

		inputData = append(inputData, input)
	}

	if err := scanner.Err(); err != nil {
		return errors.New("Error scanning standard input: " + err.Error())
	}

	cmd := exec.Command(command[0], append(command[1:], inputData...)...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

	return nil
}
