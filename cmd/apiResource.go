package cmd

import (
	"fmt"

	"github.com/operate-first/opfcli/apiresources"
	"github.com/spf13/cobra"
)

var apiResourceCmd = &cobra.Command{
	Use:   "api-resource",
	Short: "Lookup up details about an api resource",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if res, ok := apiresources.APIResourceMap[args[0]]; ok {
				fmt.Printf("%s\n", res)
			} else {
				fmt.Printf("No resource named %s\n", args[0])
			}
		} else {
			for k := range apiresources.APIResourceMap {
				fmt.Printf("%s\n", k)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(apiResourceCmd)
}
