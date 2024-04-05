package tests

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestProgramWithNormalInput(t *testing.T) {
	input := "10\n20\n30\n"
	expectedOutput := "Mean: 20.00\nMedian: 20.00\nMode: 10\nStandard Deviation: 8.16\n"

	cmd := exec.Command("../main")
	cmd.Stdin = strings.NewReader(input)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		t.Fatalf("failed to run program: %v", err)
	}

	actualOutput := stdout.String()
	if actualOutput != expectedOutput {
		t.Errorf("unexpected output: got %q, want %q", actualOutput, expectedOutput)
	}
}

func TestProgramWithOneFlag(t *testing.T) {
	input := "10\n20\n30\n"
	expectedOutput := "Mean: 20.00\n"

	cmd := exec.Command("../main", "-mean")
	cmd.Stdin = strings.NewReader(input)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		t.Fatalf("failed to run program: %v", err)
	}

	actualOutput := stdout.String()
	if actualOutput != expectedOutput {
		t.Errorf("unexpected output: got %q, want %q", actualOutput, expectedOutput)
	}
}
