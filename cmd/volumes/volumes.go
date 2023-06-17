/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// volumesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// volumesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	if !VolumesCmd.Flags().HasFlags() {
		fmt.Println("No flags defined. See help.")
	}
}
