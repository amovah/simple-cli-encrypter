/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/amovah/simple-cli-encrypter/core"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Short:   "encrypt a text",
	Aliases: []string{"enc"},
	Run: func(cmd *cobra.Command, args []string) {
		userTxt, err := prompt.New().Ask("Data:").Input("")
		if err != nil {
			log.Fatal(err)
		}

		userPass, err := prompt.New().Ask("Pasword:").Input("", input.WithEchoMode(input.EchoPassword))
		if err != nil {
			log.Fatal(err)
		}

		encrypted, err := core.Encrypt([]byte(userPass), []byte(userTxt))
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
