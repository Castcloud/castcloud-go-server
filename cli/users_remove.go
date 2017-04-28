package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Castcloud/castcloud-go-server/api"
)

var usersRemoveCmd = &cobra.Command{
	Use:   "remove <username>",
	Short: "Remove user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Usage: users remove <username>")
			return
		}

		err := api.Store().RemoveUser(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Removed", args[0])
		}
	},
}
