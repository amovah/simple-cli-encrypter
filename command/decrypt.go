/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/amovah/simple-cli-encrypter/core"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Short:   "decrypt a text",
	Aliases: []string{"dec"},
	Run: func(cmd *cobra.Command, args []string) {
		validateData := func(input string) error {
			if len(input) < 1 {
				return errors.New("text cannot be empty")
			}
			return nil
		}

		promptData := promptui.Prompt{
			Label:    "Data",
			Validate: validateData,
		}

		resultData, err := promptData.Run()
		if err != nil {
			log.Fatal(err)
		}

		validateKey := func(input string) error {
			if len(input) < 1 {
				return errors.New("key cannot be empty")
			}
			return nil
		}

		promptPassword := promptui.Prompt{
			Label:    "Password",
			Validate: validateKey,
			Mask:     '*',
		}

		resultPassword, err := promptPassword.Run()
		if err != nil {
			log.Fatal(err)
		}

		encrypted, err := hex.DecodeString(resultData)
		if err != nil {
			log.Fatal(err)
		}

		decrypted, err := core.Decrypt([]byte(resultPassword), encrypted)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("decrypted text: %s \n", decrypted)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
