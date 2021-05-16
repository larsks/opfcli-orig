package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
)

func writeKustomization(path string, resources []string, components []string) {
	kustom := models.CreateKustomization()

	if len(resources) > 0 {
		kustom.Resources = resources
	}

	if len(components) > 0 {
		kustom.Components = components
	}

	kustomOut := kustom.ToYAML()

	err := ioutil.WriteFile(
		filepath.Join(filepath.Dir(path), "kustomization.yaml"),
		[]byte(kustomOut), 0644,
	)
	if err != nil {
		log.Fatalf("failed to write kustomization: %v", err)
	}
}

func createNamespace(projectName, projectOwner, projectDescription string) {
	appName := config.GetString("app-name")
	path := filepath.Join(repoDirectory, appName, NAMESPACE_PATH, projectName, "namespaces.yaml")

	if utils.PathExists(filepath.Dir(path)) {
		log.Fatalf("namespace %s already exists", projectName)
	}

	ns := models.CreateNamespace(projectName, projectOwner, projectDescription)
	nsOut := ns.ToYAML()

	log.Printf("writing namespace definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.Fatalf("failed to create namespace directory: %v", err)
	}

	err := ioutil.WriteFile(path, []byte(nsOut), 0644)
	if err != nil {
		log.Fatalf("failed to write namespace file: %v", err)
	}

	writeKustomization(
		path,
		[]string{"namespace.yaml"},
		[]string{
			filepath.Join(COMPONENT_REL_PATH, "project-admin-rolebindings", projectOwner),
		},
	)
}

func createRoleBinding(projectName, projectOwner string) {
	appName := config.GetString("app-name")
	path := filepath.Join(repoDirectory, appName, COMPONENT_PATH, "project-admin-rolebindings", projectOwner, "rbac.yaml")

	if utils.PathExists(filepath.Dir(path)) {
		log.Printf("rolebinding already exists (continuing)")
		return
	}

	rbac := models.CreateRoleBinding(
		fmt.Sprintf("namespace-admin-%s", projectOwner),
		"admin",
	)
	rbac.Subjects = []models.Subject{
		*models.CreateGroupSubject(projectOwner),
	}
	rbacOut := rbac.ToYAML()

	log.Printf("writing rbac definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.Fatalf("failed to create rolebinding directory: %v", err)
	}

	err := ioutil.WriteFile(path, []byte(rbacOut), 0644)
	if err != nil {
		log.Fatalf("failed to write rbac: %v", err)
	}

	writeKustomization(
		path,
		[]string{"rbac.yaml"},
		nil,
	)
}

func createGroup(projectName, projectOwner string) {
	appName := config.GetString("app-name")
	path := filepath.Join(repoDirectory, appName, GROUP_PATH, projectOwner, "group.yaml")

	if utils.PathExists(filepath.Dir(path)) {
		log.Printf("group already exists (continuing)")
		return
	}

	rbac := models.CreateGroup(projectOwner)
	groupOut := rbac.ToYAML()

	log.Printf("writing group definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.Fatalf("failed to create group directory: %v", err)
	}

	err := ioutil.WriteFile(path, []byte(groupOut), 0644)
	if err != nil {
		log.Fatalf("failed to write rbac: %v", err)
	}

	writeKustomization(
		path,
		[]string{"group.yaml"},
		nil,
	)
}

var createProjectCmd = &cobra.Command{
	Use:   "create-project projectName projectOwner",
	Short: "Onboard a new project into Operate First",
	Long: `Onboard a new project into Operate First.

- Register a new group
- Register a new namespace with appropriate role bindings for your group
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectOwner := args[1]
		projectDescription := cmd.Flag("description").Value.String()

		createNamespace(projectName, projectOwner, projectDescription)
		createRoleBinding(projectName, projectOwner)
		createGroup(projectName, projectOwner)
	},
}

func init() {
	rootCmd.AddCommand(createProjectCmd)

	createProjectCmd.PersistentFlags().StringP(
		"description", "d", "", "Team description")
}
