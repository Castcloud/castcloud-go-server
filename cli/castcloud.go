package cli

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/mitchellh/go-homedir"
	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/spf13/cobra"
	"github.com/khlieng/castcloud-go/Godeps/_workspace/src/github.com/spf13/viper"

	"github.com/khlieng/castcloud-go/api"
	"github.com/khlieng/castcloud-go/assets"
)

var (
	castcloudCmd = &cobra.Command{
		Use:   os.Args[0],
		Short: "Your podcast library in the cloud.",
		Run: func(cmd *cobra.Command, args []string) {
			api.Serve()
		},
	}

	dir string
)

func init() {
	addCommands()
	bindFlags()
	setDefaults()

	cobra.OnInitialize(func() {
		dir = viper.GetString("dir")
		os.Mkdir(dir, 0777)
		initConfig()

		viper.SetConfigName("config")
		viper.AddConfigPath(dir)
		viper.ReadInConfig()

		api.Configure(&api.Config{
			Port:  viper.GetInt("port"),
			Debug: viper.GetBool("debug"),
			Dir:   dir,
			MaxDownloadConnections: viper.GetInt("crawl.max_conn"),
		})
	})
}

func Execute() {
	castcloudCmd.Execute()
}

func addCommands() {
	castcloudCmd.AddCommand(clearCmd)
	castcloudCmd.AddCommand(configCmd)
	usersCmd.AddCommand(usersAddCmd)
	usersCmd.AddCommand(usersRemoveCmd)
	castcloudCmd.AddCommand(usersCmd)
}

func bindFlags() {
	castcloudCmd.PersistentFlags().Bool("debug", false, "debug mode")
	castcloudCmd.PersistentFlags().String("dir", defaultDir(), "directory to store config and data in")
	castcloudCmd.Flags().IntP("port", "p", 3000, "port to listen on")

	viper.BindPFlag("debug", castcloudCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("dir", castcloudCmd.PersistentFlags().Lookup("dir"))
	viper.BindPFlag("port", castcloudCmd.Flags().Lookup("port"))
}

func setDefaults() {
	viper.SetDefault("crawl.max_conn", 128)
}

func initConfig() {
	configPath := path.Join(dir, "config.toml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Println("Writing default config to", configPath)

		config, err := assets.Asset("config.default.toml")
		if err != nil {
			log.Println(err)
			return
		}

		err = ioutil.WriteFile(configPath, config, 0600)
		if err != nil {
			log.Println(err)
		}
	}
}

func defaultDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(dir, ".castcloud")
}
