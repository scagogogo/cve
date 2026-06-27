package cmd

import (
	"fmt"

	cvepkg "github.com/scagogogo/cve-skills"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate CVE identifiers",
	Long:  `Generate standard format CVE identifiers.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var generateCveCmd = &cobra.Command{
	Use:   "cve --year [year] --seq [sequence]",
	Short: "Generate a CVE identifier from year and sequence number",
	Long: `Generate a standard format CVE identifier from a year and sequence number.

Examples:
  cve generate cve --year 2022 --seq 12345`,
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := cmd.Flags().GetInt("year")
		seq, _ := cmd.Flags().GetInt("seq")
		if year == 0 || seq == 0 {
			fmt.Println("error: --year and --seq are required")
			return
		}
		fmt.Println(cvepkg.GenerateCve(year, seq))
	},
}

var generateFakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "Generate a fake CVE identifier for testing",
	Long: `Generate a fake CVE identifier using the current year and a random sequence number.

This is intended for testing and demonstration purposes only.

Examples:
  cve generate fake`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cvepkg.GenerateFakeCve())
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(generateCveCmd)
	generateCmd.AddCommand(generateFakeCmd)

	generateCveCmd.Flags().IntP("year", "y", 0, "CVE year (required)")
	generateCveCmd.Flags().IntP("seq", "s", 0, "CVE sequence number (required)")
}
