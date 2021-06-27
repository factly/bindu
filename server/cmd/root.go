package cmd

import (
	"github.com/factly/bindu-server/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bindu-server",
	Short: "Backend for bindu - a data visualization platform",
	Long: `Bindu Server is a lightweight backend server for data visualization 
	platform bindu.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(config.SetupVars)
}
