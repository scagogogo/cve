package cmd

import (
	"fmt"
	"os"

	cvepkg "github.com/scagogogo/cve-skills"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate [cve-id...]",
	Short: "Validate CVE identifiers",
	Long: `Validate CVE identifiers with comprehensive checks.

Checks format, year range (1999-current year), and sequence number.

Accepts CVE identifiers as arguments or from stdin (one per line).

Examples:
  cve validate CVE-2022-12345
  echo "CVE-1998-12345" | cve validate`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			valid := cvepkg.ValidateCve(input)
			fmt.Printf("%s\t%v\n", cvepkg.Format(input), valid)
		}
	},
}

var isCveCmd = &cobra.Command{
	Use:   "is-cve [text...]",
	Short: "Check if text is exactly a CVE identifier",
	Long: `Check if the input text is exactly a valid CVE identifier format.

Returns "true" or "false" for each input.

Examples:
  cve validate is-cve CVE-2022-12345
  cve validate is-cve "text with CVE-2022-12345"`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			fmt.Printf("%s\t%v\n", input, cvepkg.IsCve(input))
		}
	},
}

var containsCveCmd = &cobra.Command{
	Use:   "contains-cve [text...]",
	Short: "Check if text contains CVE identifiers",
	Long: `Check if the input text contains any CVE identifier.

Returns "true" or "false" for each input.

Examples:
  cve validate contains-cve "System affected by CVE-2021-44228"`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			fmt.Printf("%v\n", cvepkg.IsContainsCve(input))
		}
	},
}

var yearOkCmd = &cobra.Command{
	Use:   "year-ok [cve-id...]",
	Short: "Check if CVE year is in valid range",
	Long: `Check if the CVE identifier has a valid year (1999 to current year).

Accepts an optional --cutoff flag to allow future years.

Examples:
  cve validate year-ok CVE-2022-12345
  cve validate year-ok CVE-2030-12345 --cutoff 5`,
	Run: func(cmd *cobra.Command, args []string) {
		cutoff, _ := cmd.Flags().GetInt("cutoff")
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			var result bool
			if cutoff > 0 {
				result = cvepkg.IsCveYearOkWithCutoff(input, cutoff)
			} else {
				result = cvepkg.IsCveYearOk(input)
			}
			fmt.Printf("%s\t%v\n", cvepkg.Format(input), result)
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.AddCommand(isCveCmd)
	validateCmd.AddCommand(containsCveCmd)
	validateCmd.AddCommand(yearOkCmd)
	yearOkCmd.Flags().IntP("cutoff", "c", 0, "allow N years into the future")
}
