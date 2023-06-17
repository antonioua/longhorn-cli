package cmd

import (
	"github.com/antonioua/longhorn-cli/cmd/volumes"
	"github.com/spf13/cobra"
	"os"
)

var apiUrl string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "longhorn-cli",
	Short: "This tool is used to run Longhorn commands",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true // Disable completion command

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&apiUrl, "apiUrl", "a", "http://longhorn-frontend.longhorn-system.svc.cluster.local/v1", "longhorn api url")
	//if err := rootCmd.MarkFlagRequired("apiUrl"); err != nil {
	//	fmt.Println(err)
	//}

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.longhorn-cli.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(volumes.VolumesCmd)
}
