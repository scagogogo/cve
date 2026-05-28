package cmd

import (
	"bufio"
	"os"
)

// readInputs reads CVE inputs from args or stdin.
// If args are provided, uses them directly.
// Otherwise, reads from stdin line by line.
func readInputs(args []string) []string {
	if len(args) > 0 {
		return args
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil
	}

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
