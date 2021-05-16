package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var enableMonitoringCmd = &cobra.Command{
	Use:   "enable-monitoring",
	Short: "Enable monitoring for a Kubernetes namespace",
	Long: `Enable monitoring fora Kubernetes namespace.

This will add a RoleBinding to the target namespace that permits
Prometheus to access certain metrics about pods, services, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enableMonitoring called")
	},
}

func init() {
	rootCmd.AddCommand(enableMonitoringCmd)
}
