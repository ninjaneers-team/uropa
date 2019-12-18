package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/ninjaneers-team/uropa/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	config  utils.OpaClientConfig
	verbose int
	noColor bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "urOpa",
	Short: "Administer your Opa declaritively",
	Long: `urOpa helps you manage Open Policy Agent with a declarative
configuration file.
It can be used to export, import or sync entities to Opa.`,
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if _, err := url.ParseRequestURI(config.Address); err != nil {
			return errors.WithStack(errors.Wrap(err, "invalid URL"))
		}
		if noColor {
			color.NoColor = true
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets
// sflags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var wg sync.WaitGroup
	var err error
	wg.Add(2)

	go func() {
		wg.Done()
	}()

	go func() {
		err = rootCmd.Execute()
		wg.Done()
	}()

	wg.Wait()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is $HOME/.uropa.yaml)")

	rootCmd.PersistentFlags().String("opa-addr", "http://localhost:8001",
		"HTTP Address of Opa's Admin API.\n"+
			"This value can also be set using DECK_KONG_ADDR\n"+
			" environment variable.")
	viper.BindPFlag("opa-addr",
		rootCmd.PersistentFlags().Lookup("opa-addr"))

	rootCmd.PersistentFlags().StringSlice("headers", []string{},
		"HTTP Headers(key:value) to inject in all requests to Opa's Admin API.\n"+
			"This flag can be specified multiple times to inject multiple headers.")
	viper.BindPFlag("headers",
		rootCmd.PersistentFlags().Lookup("headers"))

	rootCmd.PersistentFlags().Bool("tls-skip-verify", false,
		"Disable verification of Opa's Admin TLS certificate.\n"+
			"This value can also be set using DECK_TLS_SKIP_VERIFY "+
			"environment variable.")
	viper.BindPFlag("tls-skip-verify",
		rootCmd.PersistentFlags().Lookup("tls-skip-verify"))

	rootCmd.PersistentFlags().String("tls-server-name", "",
		"Custom CA certificate to use to verify"+
			"Opa's Admin TLS certificate.\n"+
			"This value can also be set using DECK_TLS_SERVER_NAME"+
			" environment variable.")
	viper.BindPFlag("tls-server-name",
		rootCmd.PersistentFlags().Lookup("tls-server-name"))

	rootCmd.PersistentFlags().String("ca-cert", "",
		"Custom CA certificate to use to verify"+
			"Opa's Admin TLS certificate.\n"+
			"This value can also be set using DECK_CA_CERT"+
			" environment variable.")
	viper.BindPFlag("ca-cert",
		rootCmd.PersistentFlags().Lookup("ca-cert"))

	rootCmd.PersistentFlags().Int("verbose", 0,
		"Enable verbose verbose logging levels\n"+
			"Setting this value to 2 outputs all HTTP reqeust/response\n"+
			"between urOpa and Opa.")
	viper.BindPFlag("verbose",
		rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().Bool("no-color", false,
		"disable colorized output")
	viper.BindPFlag("no-color",
		rootCmd.PersistentFlags().Lookup("no-color"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".uropa"(without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".uropa")
	}
	viper.SetEnvPrefix("uropa")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	viper.ReadInConfig()
	config.Address = viper.GetString("opa-addr")
	config.TLSServerName = viper.GetString("tls-server-name")
	config.TLSSkipVerify = viper.GetBool("tls-skip-verify")
	config.TLSCACert = viper.GetString("ca-cert")
	config.Headers = viper.GetStringSlice("headers")
	verbose = viper.GetInt("verbose")
	noColor = viper.GetBool("no-color")

	config.Debug = verbose >= 1
}
