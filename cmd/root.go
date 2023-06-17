package cmd

import (
	"fmt"
	"github.com/ddshuai/go-qywx/cmd/api"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "go-qywx",
	Short:        "go-qywx",
	SilenceUsage: true,
	Long:         `go-qywx`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	fmt.Printf("%s\n", "go qywx")
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
