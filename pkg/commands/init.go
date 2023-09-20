package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "hi",
	RunE: func(cmd *cobra.Command, args []string) error {
		println("hi! cause hi command")
		return nil
	},
}

func Exec() error {
	return rootCmd.Execute()
}
