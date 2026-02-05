package test

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestJudge(t *testing.T) {
	cmd := exec.Command("bash", "-c", "php a.php < input.in")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	outStr := stdout.String()
	errStr := stderr.String()

	if err != nil {
		fmt.Printf("Command failed: %v\n", err)
	}

	fmt.Printf("STDOUT:\n%s\n", outStr)
	fmt.Printf("STDERR:\n%s\n", errStr)
}
