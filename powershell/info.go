package powershell

import "errors"

type Domain struct {
	DomainName string `json:"DomainName"`
	DomainType string `json:"DomainType"`
	IsValid    bool   `json:"IsValid"`
}

func GetDomains() (result *[]Domain, err error) {
	powershellScript := "scripts/getdomains.ps1"
	result, err = RunExchange[[]Domain]("koksmat", powershellScript, "", "", CallbackMockup)
	if err != nil {
		return result, err
	}
	if result == nil {
		return nil, errors.New("No domains found")
	}
	return result, err

}
