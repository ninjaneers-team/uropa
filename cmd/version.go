// Copyright © 2018 Harry Bagdi <harrybagdi@gmail.com>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION is the current version of uropa.
// This should be substituted by git tag during the build process.
var VERSION = "dev"

// COMMIT is the short hash of the source tree.
// This should be substituted by Git commit hash  during the build process.
var COMMIT = "unknown"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of uropa",
	Long: `version prints the version of uropa along with git short
commit hash of the source tree`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("uropa %s (%s) \n", VERSION, COMMIT)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
