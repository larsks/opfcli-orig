package cmd

import (
	"fmt"
    "log"

	"github.com/spf13/cobra"

    "github.com/operate-first/opfcli/models"
)

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

        ns := models.CreateNamespace(projectName, projectOwner, projectDescription)
        group := models.CreateGroup(projectOwner)
        komp := models.CreateKomponent()
        rb := models.CreateRoleBinding(
            fmt.Sprintf("namespace-admin-%s", projectOwner), "admin")
        rb.Subjects = []models.Subject{
            *models.CreateGroupSubject(projectOwner),
        }

        s, err := ns.ToYAML()
        if err != nil {
            log.Fatalf("error: %v", err)
        }
        fmt.Printf("%s\n", s)

        s, err = group.ToYAML()
        if err != nil {
            log.Fatalf("error: %v", err)
        }
        fmt.Printf("%s\n", s)

        s, err = komp.ToYAML()
        if err != nil {
            log.Fatalf("error: %v", err)
        }
        fmt.Printf("%s\n", s)

        s, err = rb.ToYAML()
        if err != nil {
            log.Fatalf("error: %v", err)
        }
        fmt.Printf("%s\n", s)
	},
}

var projectDescription string
var projectOwner string
var projectName string

func init() {
	rootCmd.AddCommand(createProjectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createProjectCmd.PersistentFlags().String("foo", "", "A help for foo")
    createProjectCmd.PersistentFlags().StringVarP(
        &projectDescription, "description", "d", "", "Team description")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createProjectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
