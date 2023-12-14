/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/neckhair/bx/config"
	"github.com/neckhair/bx/internal/cli"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Creates the initial config file",
	Long:  `Run this command to create your config file. This file will hold the credentials to access the Bexio API.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Set port dynamically from settings.
		cli.Bold("💡 Bexio Credentials Setup")
		fmt.Println()
		fmt.Println("Go to https://developer.bexio.com and create a new application.")
		fmt.Println("The redirect URL should be set to http://localhost:50424.")
		fmt.Println()
		fmt.Println("Then enter the provided credentials here.")

		clientID, clientSecret := readCredentialsFromUser()
		config.SetCredentials(clientID, clientSecret)

		if err := loginAndStoreToken(cmd.Context()); err != nil {
			return fmt.Errorf("login failed: %w", err)
		}

		fmt.Println()
		fmt.Println("✅ You are all set. Other commands can now be used.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readCredentialsFromUser() (string, string) {
	clientID := cli.PromptGetInput("Client ID", "Please provide a client ID.")
	clientSecret := cli.PromptGetInput("Client Secret", "Please provide a client secret.")

	return clientID, clientSecret
}
