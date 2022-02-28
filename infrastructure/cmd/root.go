package cmd

import (
	"os"

	"github.com/cmparrela/go-clean-architecture/infrastructure/repositories/persistence"
	"github.com/cmparrela/go-clean-architecture/usecase/user"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func Execute(database *gorm.DB) {
	// dependency
	userRepostiory := persistence.NewUserRepository(database)
	userService := user.NewUserService(userRepostiory)

	// userCmd
	userCmd := NewUserCmd()
	userCmd.AddCommand(NewCreateUserCmd(userService))
	userCmd.AddCommand(NewUpdateUserCmd(userService))

	// rootCmd
	rootCmd := NewRootCmd()
	rootCmd.AddCommand(userCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-clean-arch",
		Short: "This application exemplifies the use of clean architecture using go language",
	}
	return rootCmd
}
