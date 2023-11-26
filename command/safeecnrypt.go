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

// safeecnryptCmd represents the safeecnrypt command
var safeecnryptCmd = &cobra.Command{
	Use:   "safeecnrypt",
	Short: "encrypt safely",
	Run: func(cmd *cobra.Command, args []string) {
		userTxt, err := prompt.New().Ask("Data:").Write("", write.WithCharLimit(math.MaxInt))
		if err != nil {
			log.Fatal(err)
		}

		userPass, err := prompt.New().Ask("Pasword:").Input("", input.WithEchoMode(input.EchoPassword))
		if err != nil {
			log.Fatal(err)
		}

		userPassAgain, err := prompt.New().Ask("Password Again:").Input("", input.WithEchoMode(input.EchoPassword))
		if err != nil {
			log.Fatal(err)
		}

		if userPass != userPassAgain {
			fmt.Println("password does not match")
		}

		encrypted, err := core.Encrypt([]byte(userPass), []byte(userTxt))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("encrypted text: %s \n", hex.EncodeToString(encrypted))
	},
}

func init() {
	rootCmd.AddCommand(safeecnryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// safeecnryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// safeecnryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
