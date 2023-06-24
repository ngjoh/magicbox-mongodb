package sharepoint

import (
	"embed"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/api"
	strategy "github.com/koltyakov/gosip/auth/azurecert"
	"github.com/spf13/viper"
)

//go:embed assets
var Assets embed.FS

func writeCert(filename string, b64 string) error {
	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}

func GetClient(siteUrl string) (*api.SP, error) {
	err := writeCert("pnp.pfx", viper.GetString("PNPCERTIFICATE"))
	if err != nil {
		return nil, err
	}
	auth := &strategy.AuthCnfg{
		SiteURL:  siteUrl,
		TenantID: viper.GetString("PNPTENANTID"),
		ClientID: viper.GetString("PNPAPPID"),
		CertPath: "pnp.pfx",
		CertPass: "",
	}

	client := &gosip.SPClient{AuthCnfg: auth}

	sp := api.NewSP(client)

	return sp, nil
}
func Ping(siteUrl string) error {
	sp, err := GetClient(siteUrl)
	if err != nil {
		return err
	}
	res, err := sp.Web().Select("Title").Get()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", res.Data().Title)
	return nil
}

func FilterTemplate(templateName string) error {

	_, err := Assets.ReadFile(templateName)
	if err != nil {
		return err
	}

	return nil
}
