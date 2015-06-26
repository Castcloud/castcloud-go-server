package cli

import (
	"fmt"

	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/spf13/cobra"

	"github.com/khlieng/castcloud-go/api"
)

var usersRemoveCmd = &cobra.Command{
	Use:   "remove <username>",
	Short: "Remove user",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.Store().RemoveUser(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Removed", args[0])
		}
	},
}
