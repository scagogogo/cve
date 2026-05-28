package cmd

import (
	"fmt"
	"os"

	"github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cve",
	Short: "A CLI tool for CVE identifier processing",
	Long: fmt.Sprintf(`A comprehensive CLI tool for handling CVE (Common Vulnerabilities and Exposures) identifiers.

Provides commands for formatting, validating, extracting, comparing,
sorting, filtering, grouping, and generating CVE identifiers.

Version: %s`, cve.Version),
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "suppress non-essential output")
}
