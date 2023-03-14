/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/amovah/simple-cli-encrypter/core"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Short:   "encrypt a text",
	Aliases: []string{"enc"},
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

		encrypted, err := core.Encrypt([]byte(resultPassword), []byte(resultData))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("encrypted text: %s \n", hex.EncodeToString(encrypted))
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
