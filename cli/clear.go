package cli

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear database",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Remove(path.Join(dir, "store"))
		if err == nil {
			fmt.Println("Database cleared")
		} else if os.IsNotExist(err) {
			fmt.Println("No database found")
		} else {
			fmt.Println(err)
		}
	},
}
