package cmd

import (
	"fmt"
	"os"

	"github.com/cmparrela/go-clean-architecture/usecase/user"
	"github.com/spf13/cobra"
)

func NewCreateUserCmd(userService *user.UserService) *cobra.Command {
	userDto := user.UserInputDto{}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Use this command for create a new user",
		Run: func(cmd *cobra.Command, args []string) {
			user, err := userService.Create(&userDto)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(user)
			fmt.Println("user created successfully")
		},
	}

	createCmd.Flags().StringVarP(&userDto.Name, "name", "n", "", "User name")
	createCmd.Flags().StringVarP(&userDto.Email, "email", "e", "", "User email")
	return createCmd
}
