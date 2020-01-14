// Copyright © 2018 Harry Bagdi <harrybagdi@gmail.com>

package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	syncCmdOpaStateFile string
	syncCmdParallelism  int
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use: "sync",
	Short: "Sync performs operations to get Opa's configuration " +
		"to match the state file",
	Long: `Sync command reads the state file and performs operation on Opa
to get Opa's state in sync with the input state.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return syncMain(syncCmdOpaStateFile, false, syncCmdParallelism)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if syncCmdOpaStateFile == "" {
			return errors.New("A state file with Opa's configuration " +
				"must be specified using -s/--state flag.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringVarP(&syncCmdOpaStateFile,
		"state", "s", "Opa.yaml", "file containing Opa's configuration. "+
			"Use '-' to read from stdin.")
	syncCmd.Flags().IntVar(&syncCmdParallelism, "parallelism",
		10, "Maximum number of concurrent operations")
	syncCmd.Flags().StringSliceVar(&dumpConfig.SelectorTags,
		"select-tag", []string{},
		"only entities matching tags specified via this flag are synced.\n"+
			"Multiple tags are ANDed together.")
}
