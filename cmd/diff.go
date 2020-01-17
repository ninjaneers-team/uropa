package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	diffCmdKongStateFile   string
	diffCmdParallelism     int
	diffCmdNonZeroExitCode bool
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff the current entities in Opa with the on on disks",
	Long: `Diff is like a dry run of 'uropa sync' command.
It will load entities form Opa and then perform a diff on those with
the entities present in files locally. This allows you to see the entities
that will be created or updated or deleted.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return syncMain(diffCmdKongStateFile, true, diffCmdParallelism)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if diffCmdKongStateFile == "" {
			return errors.New("A state file with Opa's configuration " +
				"must be specified using -s/--state flag.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
	diffCmd.Flags().StringVarP(&diffCmdKongStateFile,
		"state", "s", "opa.yaml", "file containing Opa's configuration. "+
			"Use '-' to read from stdin.")
	diffCmd.Flags().IntVar(&diffCmdParallelism, "parallelism",
		10, "Maximum number of concurrent operations")
	diffCmd.Flags().BoolVar(&diffCmdNonZeroExitCode, "non-zero-exit-code",
		false, "return exit code 2 if there is a diff present,\n"+
			"exit code 0 if no diff is found,\n"+
			"and exit code 1 if an error occurs.")
}
