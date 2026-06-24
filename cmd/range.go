package cmd

import (
	"fmt"
	"strings"

	"github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var parseRangeCmd = &cobra.Command{
	Use:   "parse-range <range-expr>",
	Short: "Parse a CVE range expression",
	Long:  `Parse a CVE range expression (supports "to", "..", "-" syntax) and expand into individual CVEs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (range expression)")
		}
		rangeExpr := strings.TrimSpace(inputs[0])
		result := cve.ParseCveRange(rangeExpr)
		if result == nil {
			return fmt.Errorf("invalid range expression: %s", rangeExpr)
		}
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

var isConsecutiveCmd = &cobra.Command{
	Use:   "is-consecutive <cve-a> <cve-b>",
	Short: "Check if two CVEs are consecutive",
	Long:  `Check whether two CVE identifiers are consecutive (same year, adjacent sequence numbers).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires exactly 2 CVE identifiers")
		}
		result := cve.IsCvesConsecutive(inputs[0], inputs[1])
		if result {
			fmt.Printf("%s and %s are consecutive\n", inputs[0], inputs[1])
		} else {
			fmt.Printf("%s and %s are NOT consecutive\n", inputs[0], inputs[1])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(parseRangeCmd)
	rootCmd.AddCommand(isConsecutiveCmd)
}
