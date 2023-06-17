package volumes

import (
	"fmt"
	"github.com/spf13/cobra"
)

// VolumesCmd represents the volumes command
var VolumesCmd = &cobra.Command{
	Use:   "volumes",
	Short: "Operation with Longhorn volumes",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("volumes called")
	},
}

func init() {
	if !VolumesCmd.Flags().HasFlags() {
		fmt.Println("No flags defined. See help.")
	}
}
