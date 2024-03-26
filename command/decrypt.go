/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package command

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"

	"github.com/amovah/simple-cli-encrypter/core"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
	"github.com/cqroot/prompt/write"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Short:   "decrypt a text",
	Aliases: []string{"dec"},
	Run: func(cmd *cobra.Command, args []string) {
		userTxt, err := prompt.New().Ask("Data:").Write("", write.WithCharLimit(math.MaxInt))
		if err != nil {
			log.Fatal(err)
		}

		userPass, err := prompt.New().Ask("Pasword:").Input("", input.WithEchoMode(input.EchoPassword))
		if err != nil {
			log.Fatal(err)
		}

		encrypted, err := hex.DecodeString(userTxt)
		if err != nil {
			log.Fatal(err)
		}

		decrypted, err := core.Decrypt([]byte(userPass), encrypted)
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
