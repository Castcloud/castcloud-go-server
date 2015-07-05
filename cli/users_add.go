package cli

import (
	"fmt"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/howeyc/gopass"
	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/spf13/cobra"

	"github.com/Castcloud/castcloud-go-server/api"

	. "github.com/Castcloud/castcloud-go-server/api/schema"
)

var usersAddCmd = &cobra.Command{
	Use:   "add <username>",
	Short: "Add new user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Usage: users add <username>")
			return
		}

		user := api.Store().GetUser(args[0])
		if user != nil {
			fmt.Println("Username already in use")
			return
		}

		user = &User{
			Username: args[0],
			Password: getPassword(),
		}

		err := api.Store().AddUser(user)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Added", user.Username)
		}
	},
}

func getPassword() string {
	for {
		fmt.Print("Enter password: ")
		pass := string(gopass.GetPasswdMasked())
		fmt.Print("Re-enter password: ")
		passVerify := string(gopass.GetPasswdMasked())

		if pass == passVerify {
			return pass
		}

		fmt.Println("Passwords do not match, try again")
	}
}
