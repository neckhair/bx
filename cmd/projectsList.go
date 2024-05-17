package cmd

import (
	"fmt"

	"github.com/neckhair/bx/bexio"
	"github.com/neckhair/bx/config"
	"github.com/neckhair/bx/internal/cli"
	"github.com/spf13/cobra"
)

var projectListHeaders = []string{"ID", "NAME", "STATE"}

// projectsListCmd represents the projectsList command
var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		tokenSource, err := config.NewCachingTokenSource()
		if err != nil {
			return err
		}
		client := bexio.NewClient(cmd.Context(), tokenSource)

		s := cli.StartSpinner("Loading projects...")

		projects, err := client.GetProjectsByState(bexio.ProjectStateActive)
		if err != nil {
			return err
		}

		s.Stop()

		rows := make([][]string, len(projects))
		for i, project := range projects {
			rows[i] = []string{
				fmt.Sprintf("%8d", project.ID),
				project.Name,
				project.StateName(),
			}
		}

		cli.PrintTable(projectListHeaders, rows)

		return nil
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}
