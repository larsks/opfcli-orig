package cmd

import (
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/operate-first/opfcli/utils"
)

func addMonitoringRBAC(projectName string) {
	appName := config.GetString("app-name")
	path := filepath.Join(repoDirectory, appName, NAMESPACE_PATH, projectName)

	if !utils.PathExists(path) {
		log.Fatalf("namespace %s does not exist", projectName)
	}

	log.Printf("enabling monitoring on namespace %s", projectName)
	kustomize := exec.Command(
		"kustomize", "edit", "add", "component",
		filepath.Join(COMPONENT_REL_PATH, "monitoring-rbac"),
	)
	kustomize.Dir = path

	// NB: if the specified component does not exist, kustomize will fail to
	// edit the file and emit a log message but will not return an error code.
	if err := kustomize.Run(); err != nil {
		log.Fatalf("failed to edit kustomization: %v", err)
	}
}

var enableMonitoringCmd = &cobra.Command{
	Use:   "enable-monitoring",
	Short: "Enable monitoring for a Kubernetes namespace",
	Long: `Enable monitoring fora Kubernetes namespace.

This will add a RoleBinding to the target namespace that permits
Prometheus to access certain metrics about pods, services, etc.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		addMonitoringRBAC(args[0])
	},
}

func init() {
	rootCmd.AddCommand(enableMonitoringCmd)
}
