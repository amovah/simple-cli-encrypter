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

// dfileCmd represents the dfile command
var dfileCmd = &cobra.Command{
	Use:   "dfile",
	Short: "decrypt a file",
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

		encryptedTxt := make([]byte, hex.DecodedLen(len(fileContent)))
		_, err = hex.Decode(encryptedTxt, fileContent)
		if err != nil {
			log.Fatal(err)
		}

		result, err := core.Decrypt([]byte(userPass), encryptedTxt)
		if err != nil {
			log.Fatal(err)
		}

		outputPath, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatal(err)
		}

		if len(outputPath) > 0 {
			err := os.WriteFile(outputPath, []byte(result), 0777)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Printf("decrypted text: %s \n", result)
		}
	},
}

func init() {
	rootCmd.AddCommand(dfileCmd)

	dfileCmd.Flags().StringP("file", "f", "", "file path")
	dfileCmd.MarkFlagRequired("file")
	dfileCmd.Flags().StringP("output", "o", "", "store result into a file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
