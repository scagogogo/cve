package cmd

import (
	"fmt"

	"github.com/scagogogo/cve-skills"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the version number of the cve tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cve.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
