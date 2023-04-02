/*
Copyright © 2023 mekka <chmikata@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/chmikata/csvconvert/internal/logic"
	"github.com/spf13/cobra"
)

var outputFile string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create CSV from contract-data and detail-data",
	Long:  "create CSV from contract-data and detail-data.",
	Run: func(cmd *cobra.Command, args []string) {
		c := logic.NewController(
			logic.NewContractCSVCreater(),
			logic.NewFileWriter(logic.WithFilePath(outputFile)),
		)
		if err := c.CreateCSV(); err != nil {
			fmt.Println("CSV create failed.")
			os.Exit(1)
		}
		fmt.Println("CSV create success.")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "out.csv", "output file path.")
}
