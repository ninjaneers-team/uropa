// Copyright © 2018 Harry Bagdi <harrybagdi@gmail.com>

package cmd

import (
	"strings"

	"github.com/ninjaneers-team/uropa/dump"
	"github.com/ninjaneers-team/uropa/file"
	"github.com/ninjaneers-team/uropa/state"
	"github.com/ninjaneers-team/uropa/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	dumpCmdOpaStateFile string
	dumpCmdStateFormat  string
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Export Opa configuration to a file",
	Long: `Dump command reads all the entities present in Opa
and writes them to a file on disk.
The file can then be read using the Sync o Diff command to again
configure Opa.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		client, err := utils.GetOpaClient(config)
		if err != nil {
			return err
		}

		format := file.Format(strings.ToUpper(dumpCmdStateFormat))

		rawState, err := dump.Get(client, dumpConfig)
		if err != nil {
			return errors.Wrap(err, "reading configuration from Opa")
		}
		ks, err := state.Get(rawState)
		if err != nil {
			return errors.Wrap(err, "building state")
		}
		if err := file.OpaStateToFile(ks, file.WriteConfig{
			Filename:   dumpCmdOpaStateFile,
			FileFormat: format,
		}); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.Flags().StringVarP(&dumpCmdOpaStateFile, "output-file", "o",
		"opa", "file to which to write Opa's configuration."+
			"Use '-' to write to stdout.")
	dumpCmd.Flags().StringVar(&dumpCmdStateFormat, "format",
		"yaml", "output file format: json or yaml")

}
