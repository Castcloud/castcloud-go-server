package cli

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/spf13/cobra"
	"github.com/xenolf/lego/acme"
)

const caURL = "https://acme-v01.api.letsencrypt.org/directory"
const rsaKeySize = 2048

var sslCmd = &cobra.Command{
	Use:   "ssl <domain>",
	Short: "Enable SSL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Usage: ssl <domain>")
			return
		}

		sslDir := path.Join(dir, "ssl")
		err := os.Mkdir(sslDir, 0777)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}

		letsEncrypt(args[0], sslDir)
	},
}

type acmeUser struct {
	Registration *acme.RegistrationResource
	key          *rsa.PrivateKey
}

func (u acmeUser) GetEmail() string {
	return ""
}

func (u acmeUser) GetRegistration() *acme.RegistrationResource {
	return u.Registration
}

func (u acmeUser) GetPrivateKey() *rsa.PrivateKey {
	return u.key
}

func letsEncrypt(domain, outputDir string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		log.Fatal(err)
	}

	user := acmeUser{
		key: privateKey,
	}

	client, err := acme.NewClient(caURL, &user, rsaKeySize, "443")
	if err != nil {
		log.Fatal(err)
	}

	reg, err := client.Register()
	if err != nil {
		log.Fatal(err)
	}
	user.Registration = reg

	err = client.AgreeToTOS()
	if err != nil {
		log.Fatal(err)
	}

	certs, errors := client.ObtainCertificates([]string{domain}, false)
	if len(errors) > 0 {
		for k, err := range errors {
			log.Println(k, err)
		}

		os.Exit(1)
	}

	ioutil.WriteFile(path.Join(outputDir, "cert"), certs[0].Certificate, 0777)
	ioutil.WriteFile(path.Join(outputDir, "key"), certs[0].PrivateKey, 0777)
}
