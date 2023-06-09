package magicbox

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func Powerpack(payloadBody any) (result []byte, err error) {
	url := viper.GetString("MAGICBOXURL")

	method := "POST"

	client := &http.Client{}
	JSON, err := json.Marshal(payloadBody)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(JSON))

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if (res.StatusCode != 200) && (res.StatusCode != 201) {
		body, _ := ioutil.ReadAll(res.Body)
		newErr := errors.New(string(body))
		return nil, newErr
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)

}
