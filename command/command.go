package command

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: "app"}
)

func Execute() error {
	return rootCmd.Execute()
}
