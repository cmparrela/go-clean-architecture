package cmd

import "github.com/spf13/cobra"

func NewUserCmd() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Use this command to manipulate user data",
	}
	return userCmd
}
