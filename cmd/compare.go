package cmd

import (
	"fmt"
	"os"

	cvepkg "github.com/scagogogo/cve-skills"
	"github.com/spf13/cobra"
)

var compareCmd = &cobra.Command{
	Use:   "compare [cve-a] [cve-b]",
	Short: "Compare two CVE identifiers",
	Long: `Compare two CVE identifiers by year and sequence number.

Output: -1 (a < b), 0 (a == b), 1 (a > b)

Examples:
  cve compare CVE-2021-44228 CVE-2022-12345`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result := cvepkg.CompareCves(args[0], args[1])
		fmt.Println(result)
	},
}

var sortCmd = &cobra.Command{
	Use:   "sort [cve-id...]",
	Short: "Sort CVE identifiers by year and sequence number",
	Long: `Sort CVE identifiers in ascending order by year and sequence number.

Accepts CVE identifiers as arguments or from stdin (one per line).

Examples:
  cve compare sort CVE-2022-2222 CVE-2020-1111 CVE-2022-1111`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		sorted := cvepkg.SortCves(inputs)
		for _, c := range sorted {
			fmt.Println(c)
		}
	},
}

var compareByYearCmd = &cobra.Command{
	Use:   "by-year [cve-a] [cve-b]",
	Short: "Compare two CVE identifiers by year only",
	Long: `Compare two CVE identifiers by year only.

Output: negative (a earlier), 0 (same year), positive (a later)

Examples:
  cve compare by-year CVE-2021-44228 CVE-2022-12345`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result := cvepkg.CompareByYear(args[0], args[1])
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
	compareCmd.AddCommand(sortCmd)
	compareCmd.AddCommand(compareByYearCmd)
}
