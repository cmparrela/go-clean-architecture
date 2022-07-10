package cmd

import (
	"fmt"
	"os"

	"github.com/cmparrela/go-clean-architecture/usecase/user"
	"github.com/spf13/cobra"
)

func NewUpdateUserCmd(userService user.Service) *cobra.Command {
	var id uint
	userDto := user.CreateDto{}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Use this command for update a user",
		Run: func(cmd *cobra.Command, args []string) {
			user, err := userService.Update(id, &userDto)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(user)
			fmt.Println("user updated successfully")
		},
	}

	updateCmd.Flags().UintVarP(&id, "id", "i", 0, "User ID")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVarP(&userDto.Name, "name", "n", "", "User name")
	updateCmd.Flags().StringVarP(&userDto.Email, "email", "e", "", "User email")

	return updateCmd
}
