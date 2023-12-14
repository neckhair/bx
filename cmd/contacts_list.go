package cmd

import (
	"strings"

	"github.com/neckhair/bx/bexio"
	"github.com/neckhair/bx/config"
	"github.com/neckhair/bx/internal/cli"
	"github.com/spf13/cobra"
)

const contactsDefaultLimit = 500

var contactListHeaders = []string{"ID", "NAME", "ADDRESS", "POSTCODE", "CITY"}

func sanitizeString(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(str), "\n\r", "\n"), "\r", "")
}

// contactsListCmd represents the contactsList command
var contactsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List contacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := bexio.NewClient(config.TokenProvider)

		limit, err := cmd.Flags().GetInt("limit")
		if err != nil {
			return err
		}

		contacts, err := client.Contacts(cmd.Context(), limit)
		if err != nil {
			return err
		}

		rows := make([][]string, len(contacts))
		for i, contact := range contacts {
			rows[i] = []string{
				contact.Number,
				sanitizeString(contact.FullName()),
				sanitizeString(contact.Address),
				contact.Postcode,
				sanitizeString(contact.City),
			}
		}

		cli.PrintTable(contactListHeaders, rows)

		return nil
	},
}

func init() {
	contactsCmd.AddCommand(contactsListCmd)
	contactsListCmd.PersistentFlags().Int("limit", contactsDefaultLimit, "Maximum number of records to show")
}
