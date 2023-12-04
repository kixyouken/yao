package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sskCmd = &cobra.Command{
	Use:   "ssk",
	Short: L("Show ssk"),
	Long:  L("Show ssk"),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("这是我添加的SSK...")
	},
}
