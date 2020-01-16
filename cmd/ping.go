package cmd

import (
	"fmt"

	"github.com/ninjaneers-team/uropa/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Verify connectivity with Opa",
	Long: `Ping command can be used to verify if uropa
can connect to Opa's HTTP API or not.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := utils.GetOpaClient(config)
		if err != nil {
			return errors.Wrap(err, "creating opa client")
		}
		_, err = client.Health(nil)
		if err != nil {
			return errors.Wrap(err, "connecting to opa")
		}
		fmt.Println("Successfully connected to Open Policy Agent!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
