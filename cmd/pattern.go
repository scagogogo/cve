package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scagogogo/cve-skills"
	"github.com/spf13/cobra"
)

var filterPatternCmd = &cobra.Command{
	Use:   "filter-pattern <pattern> <cve-list>",
	Short: "Filter CVEs by wildcard pattern",
	Long:  `Filter CVE identifiers using wildcard pattern matching (e.g., "CVE-2022-*").`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires pattern and CVE list")
		}
		pattern := strings.TrimSpace(inputs[0])
		var cveList []string
		for _, input := range inputs[1:] {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		result := cve.FilterCvesByPattern(cveList, pattern)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

var formatSeqCmd = &cobra.Command{
	Use:   "format-seq <width> <cve>",
	Short: "Format CVE sequence number with zero-padding",
	Long:  `Format a CVE's sequence number to a fixed width with leading zeros.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires width and CVE identifier")
		}
		width, err := strconv.Atoi(strings.TrimSpace(inputs[0]))
		if err != nil {
			return fmt.Errorf("invalid width: %s", inputs[0])
		}
		result := cve.FormatSeq(inputs[1], width)
		fmt.Println(result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(filterPatternCmd)
	rootCmd.AddCommand(formatSeqCmd)
}
