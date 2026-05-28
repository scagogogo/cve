package cmd

import (
	"fmt"
	"os"

	cvepkg "github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract [text...]",
	Short: "Extract CVE identifiers from text",
	Long: `Extract CVE identifiers from text.

By default extracts all CVE identifiers found in the input text.

Accepts text as arguments or from stdin.

Examples:
  cve extract "System affected by CVE-2021-44228 and CVE-2022-12345"
  echo "CVE-2021-44228 vulnerability" | cve extract`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			cves := cvepkg.ExtractCve(input)
			for _, c := range cves {
				fmt.Println(c)
			}
		}
	},
}

var extractFirstCmd = &cobra.Command{
	Use:   "first [text...]",
	Short: "Extract the first CVE identifier from text",
	Long: `Extract the first CVE identifier found in the input text.

Examples:
  cve extract first "CVE-2021-44228 and CVE-2022-12345"`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			fmt.Println(cvepkg.ExtractFirstCve(input))
		}
	},
}

var extractLastCmd = &cobra.Command{
	Use:   "last [text...]",
	Short: "Extract the last CVE identifier from text",
	Long: `Extract the last CVE identifier found in the input text.

Examples:
  cve extract last "CVE-2021-44228 and CVE-2022-12345"`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			fmt.Println(cvepkg.ExtractLastCve(input))
		}
	},
}

var extractYearCmd = &cobra.Command{
	Use:   "year [cve-id...]",
	Short: "Extract the year from a CVE identifier",
	Long: `Extract the year part from CVE identifiers.

Examples:
  cve extract year CVE-2022-12345`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			fmt.Println(cvepkg.ExtractCveYear(input))
		}
	},
}

var extractSeqCmd = &cobra.Command{
	Use:   "seq [cve-id...]",
	Short: "Extract the sequence number from a CVE identifier",
	Long: `Extract the sequence number part from CVE identifiers.

Examples:
  cve extract seq CVE-2022-12345`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			fmt.Println(cvepkg.ExtractCveSeq(input))
		}
	},
}

var extractSplitCmd = &cobra.Command{
	Use:   "split [cve-id...]",
	Short: "Split CVE identifier into year and sequence number",
	Long: `Split CVE identifiers into their year and sequence number components.

Output format: year<TAB>sequence

Examples:
  cve extract split CVE-2022-12345`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			year, seq := cvepkg.Split(input)
			fmt.Printf("%s\t%s\n", year, seq)
		}
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)
	extractCmd.AddCommand(extractFirstCmd)
	extractCmd.AddCommand(extractLastCmd)
	extractCmd.AddCommand(extractYearCmd)
	extractCmd.AddCommand(extractSeqCmd)
	extractCmd.AddCommand(extractSplitCmd)
}
