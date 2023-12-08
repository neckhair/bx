package cmd

import (
	"github.com/spf13/cobra"
)

// contactsCmd represents the contact command
var contactsCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Manage contacts",
}

func init() {
	rootCmd.AddCommand(contactsCmd)
}
