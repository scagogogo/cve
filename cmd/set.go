package cmd

import (
	"fmt"
	"strings"

	"github.com/scagogogo/cve-skills"
	"github.com/spf13/cobra"
)

var intersectCmd = &cobra.Command{
	Use:   "intersect <list1> <list2>",
	Short: "Compute intersection of two CVE lists",
	Long:  `Compute the intersection (CVEs in both lists) of two CVE lists. Input lists may be comma-separated or read from stdin.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires exactly 2 arguments (two CVE lists)")
		}
		list1 := strings.Split(inputs[0], ",")
		list2 := strings.Split(inputs[1], ",")
		result := cve.IntersectCves(list1, list2)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

var unionCmd = &cobra.Command{
	Use:   "union <list1> <list2>",
	Short: "Compute union of two CVE lists",
	Long:  `Compute the union (all CVEs from both lists) of two CVE lists.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires exactly 2 arguments (two CVE lists)")
		}
		list1 := strings.Split(inputs[0], ",")
		list2 := strings.Split(inputs[1], ",")
		result := cve.UnionCves(list1, list2)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

var diffCmd = &cobra.Command{
	Use:   "diff <list1> <list2>",
	Short: "Compute difference (a - b) of two CVE lists",
	Long:  `Compute the difference (CVEs in list1 but not in list2) of two CVE lists.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		inputs := readInputs(args)
		if len(inputs) < 2 {
			return fmt.Errorf("requires exactly 2 arguments (two CVE lists)")
		}
		list1 := strings.Split(inputs[0], ",")
		list2 := strings.Split(inputs[1], ",")
		result := cve.DiffCves(list1, list2)
		for _, v := range result {
			fmt.Println(v)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(intersectCmd)
	rootCmd.AddCommand(unionCmd)
	rootCmd.AddCommand(diffCmd)
}
