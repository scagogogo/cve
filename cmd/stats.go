package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var countByYearCmd = &cobra.Command{
	Use:   "count-by-year <cve-list>",
	Short: "Count CVEs by year",
	Long:  `Group and count CVE identifiers by year.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (CVE list)")
		}
		var cveList []string
		for _, input := range inputs {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		counts := cve.CountByYear(cveList)
		for year, count := range counts {
			fmt.Printf("%d: %d\n", year, count)
		}
		return nil
	},
}

var yearRangeCmd = &cobra.Command{
	Use:   "year-range <cve-list>",
	Short: "Get the earliest and latest year of CVEs",
	Long:  `Get the earliest and latest year from a CVE list.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			return fmt.Errorf("requires at least 1 argument (CVE list)")
		}
		var cveList []string
		for _, input := range inputs {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		min, max := cve.YearRange(cveList)
		fmt.Printf("Year range: %d - %d (span: %d years)\n", min, max, max-min)
		return nil
	},
}

var seqRangeCmd = &cobra.Command{
	Use:   "seq-range <year> <cve-list>",
	Short: "Get sequence number range for a given year",
	Long:  `Get the smallest and largest sequence number for CVEs in a specific year.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires year and CVE list")
		}
		year, err := strconv.Atoi(strings.TrimSpace(inputs[0]))
		if err != nil {
			return fmt.Errorf("invalid year: %s", inputs[0])
		}
		var cveList []string
		for _, input := range inputs[1:] {
			cveList = append(cveList, strings.Split(input, ",")...)
		}
		min, max := cve.SeqRange(cveList, year)
		fmt.Printf("Year %d sequence range: %d - %d\n", year, min, max)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(countByYearCmd)
	rootCmd.AddCommand(yearRangeCmd)
	rootCmd.AddCommand(seqRangeCmd)
}
