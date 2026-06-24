package cmd

import (
	"fmt"
	"strings"

	"github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var validateBatchCmd = &cobra.Command{
	Use:   "validate-batch <cve-list>",
	Short: "Batch validate CVE identifiers",
	Long:  `Validate a batch of CVE identifiers and report detailed results for each.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (CVE list)")
		}
		var cveList []string
		for _, input := range inputs {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		results := cve.ValidateCves(cveList)
		for _, r := range results {
			if r.Valid {
				fmt.Printf("✓ %s\n", r.Cve)
			} else {
				fmt.Printf("✗ %s — %s\n", r.Cve, r.Reason)
			}
		}
		return nil
	},
}

var filterValidCmd = &cobra.Command{
	Use:   "filter-valid <cve-list>",
	Short: "Filter out only valid CVE identifiers",
	Long:  `Filter a list of CVE identifiers, keeping only valid ones.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (CVE list)")
		}
		var cveList []string
		for _, input := range inputs {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		result := cve.FilterValidCves(cveList)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateBatchCmd)
	rootCmd.AddCommand(filterValidCmd)
}
