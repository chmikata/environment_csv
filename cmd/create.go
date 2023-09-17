/*
Copyright © 2023 mekka <chmikata@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/chmikata/environment_csv/internal/application"
	"github.com/chmikata/environment_csv/internal/infra"
	"github.com/chmikata/environment_csv/internal/userinterface/implement"
	"github.com/spf13/cobra"
)

var outputFile string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create environment CSV from contract-data and detail-data",
	Long:  "create environment CSV from contract-data and detail-data.",
	Run: func(cmd *cobra.Command, args []string) {
		c := application.NewController(
			infra.NewHbEnvironment("localhost", 5432,
				infra.WithUser("postgres"),
				infra.WithPass("postgres"),
				infra.WithDbName("environment"),
			),
			implement.NewFileWriter(implement.WithFilePath(outputFile)),
		)
		if err := c.CreateEnvironment(); err != nil {
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
