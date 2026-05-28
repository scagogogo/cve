package cmd

import (
	"fmt"
	"os"
	"sort"

	cvepkg "github.com/scagogogo/cve"
	"github.com/spf13/cobra"
)

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filter and group CVE identifiers",
	Long: `Filter and group CVE identifiers by various criteria.

Subcommands: by-year, by-year-range, recent, group-by-year, dedup

Accepts CVE identifiers as arguments or from stdin (one per line).`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var filterByYearCmd = &cobra.Command{
	Use:   "by-year --year [year] [cve-id...]",
	Short: "Filter CVEs by specific year",
	Long: `Filter CVE identifiers by a specific year.

Examples:
  cve filter by-year --year 2022 CVE-2021-1111 CVE-2022-2222
  echo -e "CVE-2021-1111\nCVE-2022-2222" | cve filter by-year --year 2022`,
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := cmd.Flags().GetInt("year")
		if year == 0 {
			fmt.Fprintln(os.Stderr, "error: --year is required")
			os.Exit(1)
		}
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		filtered := cvepkg.FilterCvesByYear(inputs, year)
		for _, c := range filtered {
			fmt.Println(c)
		}
	},
}

var filterByYearRangeCmd = &cobra.Command{
	Use:   "by-year-range --start [year] --end [year] [cve-id...]",
	Short: "Filter CVEs by year range",
	Long: `Filter CVE identifiers within a year range (inclusive).

Examples:
  cve filter by-year-range --start 2021 --end 2022 CVE-2020-1111 CVE-2021-2222 CVE-2022-3333`,
	Run: func(cmd *cobra.Command, args []string) {
		startYear, _ := cmd.Flags().GetInt("start")
		endYear, _ := cmd.Flags().GetInt("end")
		if startYear == 0 || endYear == 0 {
			fmt.Fprintln(os.Stderr, "error: --start and --end are required")
			os.Exit(1)
		}
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		filtered := cvepkg.FilterCvesByYearRange(inputs, startYear, endYear)
		for _, c := range filtered {
			fmt.Println(c)
		}
	},
}

var filterRecentCmd = &cobra.Command{
	Use:   "recent --years [n] [cve-id...]",
	Short: "Filter CVEs from recent N years",
	Long: `Filter CVE identifiers from the most recent N years.

Examples:
  cve filter recent --years 2 CVE-2020-1111 CVE-2022-2222 CVE-2023-3333`,
	Run: func(cmd *cobra.Command, args []string) {
		years, _ := cmd.Flags().GetInt("years")
		if years == 0 {
			fmt.Fprintln(os.Stderr, "error: --years is required")
			os.Exit(1)
		}
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		filtered := cvepkg.GetRecentCves(inputs, years)
		for _, c := range filtered {
			fmt.Println(c)
		}
	},
}

var groupByYearCmd = &cobra.Command{
	Use:   "group-by-year [cve-id...]",
	Short: "Group CVEs by year",
	Long: `Group CVE identifiers by year.

Output format: year: then one CVE per indented line.

Examples:
  cve filter group-by-year CVE-2021-1111 CVE-2022-2222 CVE-2021-3333`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		groups := cvepkg.GroupByYear(inputs)
		var years []string
		for y := range groups {
			years = append(years, y)
		}
		sort.Strings(years)
		for _, y := range years {
			fmt.Printf("%s:\n", y)
			for _, c := range groups[y] {
				fmt.Printf("  %s\n", c)
			}
		}
	},
}

var dedupCmd = &cobra.Command{
	Use:   "dedup [cve-id...]",
	Short: "Remove duplicate CVE identifiers",
	Long: `Remove duplicate CVE identifiers (case-insensitive).

Examples:
  cve filter dedup CVE-2022-1111 cve-2022-1111 CVE-2022-2222`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		unique := cvepkg.RemoveDuplicateCves(inputs)
		for _, c := range unique {
			fmt.Println(c)
		}
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
	filterCmd.AddCommand(filterByYearCmd)
	filterCmd.AddCommand(filterByYearRangeCmd)
	filterCmd.AddCommand(filterRecentCmd)
	filterCmd.AddCommand(groupByYearCmd)
	filterCmd.AddCommand(dedupCmd)

	filterByYearCmd.Flags().IntP("year", "y", 0, "target year (required)")
	filterByYearRangeCmd.Flags().IntP("start", "s", 0, "start year (required, inclusive)")
	filterByYearRangeCmd.Flags().IntP("end", "e", 0, "end year (required, inclusive)")
	filterRecentCmd.Flags().IntP("years", "n", 0, "number of recent years (required)")
}
