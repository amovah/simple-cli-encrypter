package command

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/amovah/simple-cli-encrypter/core"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
	"github.com/spf13/cobra"
)

// efileCmd represents the efile command
var efileCmd = &cobra.Command{
	Use:   "efile",
	Short: "encrypt a file",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatal(err)
		}

		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		userPass, err := prompt.New().Ask("Pasword:").Input("", input.WithEchoMode(input.EchoPassword))
		if err != nil {
			log.Fatal(err)
		}

		encrypted, err := core.Encrypt([]byte(userPass), fileContent)
		if err != nil {
			log.Fatal(err)
		}

		result := make([]byte, hex.EncodedLen(len(encrypted)))
		hex.Encode(result, encrypted)

		outputPath, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatal(err)
		}

		if len(outputPath) > 0 {
			err := os.WriteFile(outputPath, result, 0777)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("encrypted text: %s \n", result)
		}
	},
}

func init() {
	rootCmd.AddCommand(efileCmd)

	efileCmd.Flags().StringP("file", "f", "", "file path")
	efileCmd.MarkFlagRequired("file")
	efileCmd.Flags().StringP("output", "o", "", "store result into a file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// efileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// efileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
