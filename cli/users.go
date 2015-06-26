package cli

import (
	"fmt"

	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/spf13/cobra"

	"github.com/khlieng/castcloud-go/api"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "List users",
	Run: func(cmd *cobra.Command, args []string) {
		users := api.Store().GetUsers()

		if len(users) == 0 {
			fmt.Println("There are no users")
		} else if len(users) == 1 {
			fmt.Println("1 user:")
			fmt.Println(users[0].Username)
		} else {
			fmt.Println(len(users), "users:")
			for _, user := range users {
				fmt.Println(user.Username)
			}
		}
	},
}
