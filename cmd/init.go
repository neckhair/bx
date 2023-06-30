package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/neckhair/bx/config"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the configuration",
	Long:  `Initialize the configuration and login to Bexio.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go to https://developer.bexio.com/ to create an app. Then enter the credentials here.")

		fmt.Println("Client ID:")
		clientID, err := readStringFromInput()
		if err != nil {
			log.Fatal(err)
		} else {
			config.Set("client_id", clientID)
		}

		fmt.Println("Client Secret:")
		clientSecret, err := readStringFromInput()
		if err != nil {
			log.Fatal(err)
		} else {
			config.Set("client_secret", clientSecret)
		}

		if err := config.Write(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readStringFromInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil
}
