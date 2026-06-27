package cmd

import (
	"fmt"
	"os"

	cvepkg "github.com/scagogogo/cve-skills"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format [cve-id...]",
	Short: "Format CVE identifiers to standard uppercase",
	Long: `Format CVE identifiers to standard uppercase format.

Accepts CVE identifiers as arguments or from stdin (one per line).

Examples:
  cve format CVE-2022-12345
  cve format cve-2022-12345
  echo "cve-2022-12345" | cve format`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs := readInputs(args)
		if len(inputs) == 0 {
			os.Exit(1)
		}
		for _, input := range inputs {
			fmt.Println(cvepkg.Format(input))
		}
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
