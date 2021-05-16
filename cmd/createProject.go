package cmd

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"

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

    kustom_out := kustom.ToYAML()

    err := ioutil.WriteFile(
        filepath.Join(filepath.Dir(path), "kustomization.yaml"),
        []byte(kustom_out), 0644,
    )
    if err != nil {
        log.Fatalf("failed to write kustomization: %v", err)
    }
}

func createNamespace() {
    appName := config.GetString("app-name")
    path := filepath.Join(repoDirectory, appName, NAMESPACE_PATH, projectName, "namespaces.yaml")

    if utils.PathExists(filepath.Dir(path)) {
        log.Fatalf("namespace %s already exists", projectName)
    }

    ns := models.CreateNamespace(projectName, projectOwner, projectDescription)
    ns_out := ns.ToYAML()

    log.Printf("writing namespace definition to %s", filepath.Dir(path))
    if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
        log.Fatalf("failed to create namespace directory: %v", err)
    }

    err := ioutil.WriteFile(path, []byte(ns_out), 0644)
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

func createRoleBinding() {
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
    rbac_out := rbac.ToYAML()

    log.Printf("writing rbac definition to %s", filepath.Dir(path))
    if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
        log.Fatalf("failed to create rolebinding directory: %v", err)
    }

    err := ioutil.WriteFile(path, []byte(rbac_out), 0644)
    if err != nil {
        log.Fatalf("failed to write rbac: %v", err)
    }

    writeKustomization(
        path,
        []string{"rbac.yaml"},
        nil,
    )
}

func createGroup() {
    appName := config.GetString("app-name")
    path := filepath.Join(repoDirectory, appName, GROUP_PATH, projectOwner, "group.yaml")

    if utils.PathExists(filepath.Dir(path)) {
        log.Printf("group already exists (continuing)")
        return
    }

    rbac := models.CreateGroup(projectOwner)
    group_out := rbac.ToYAML()

    log.Printf("writing group definition to %s", filepath.Dir(path))
    if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
        log.Fatalf("failed to create group directory: %v", err)
    }

    err := ioutil.WriteFile(path, []byte(group_out), 0644)
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
        projectName = args[0]
        projectOwner = args[1]

        createNamespace()
        createRoleBinding()
        createGroup()
	},
}

var projectDescription string
var projectOwner string
var projectName string

func init() {
	rootCmd.AddCommand(createProjectCmd)

    createProjectCmd.PersistentFlags().StringVarP(
        &projectDescription, "description", "d", "", "Team description")
}
