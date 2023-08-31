package powershell

import (
	"fmt"
	"strings"
)

type NewSharedMailboxResult struct {
	Name               string `json:"Name"`
	DisplayName        string `json:"DisplayName"`
	ExchangeObjectId   string `json:"ExchangeObjectId"`
	PrimarySmtpAddress string `json:"PrimarySmtpAddress"`
}
type EmptyResult struct {
}

type Member struct {
	User         string `json:"User"`
	AccessRights string `json:"AccessRights"`
	IsInherited  bool   `json:"IsInherited"`
}
type MembersResponse struct {
	Members []Member `json:"Members"`
}

type SyncResponse struct {
	Changes []string `json:"changes"`
}
type OwnersResponse struct {
	Owners string `json:"Owners"`
}

func PwshArray(members []string) string {

	pwshArray := strings.Join(members, ",")
	if (pwshArray) == "" {

		return "\"\""
	} else {
		return pwshArray
	}
}
func CreateSharedMailbox(appid string, Name string, DisplayName string, Alias string, Owners []string, Members []string, Readers []string) (result *NewSharedMailboxResult, err error) {
	powershellScript := "scripts/sharedmailboxes/create.ps1"
	powershellArguments := fmt.Sprintf(` -Name "%s" -DisplayName "%s"  -Alias "%s" -Members %s -Readers %s -Owners="%s"`, Name, DisplayName, Alias, PwshArray(Members), PwshArray(Readers), PwshArray(Owners))
	result, err = RunExchange[NewSharedMailboxResult](appid, powershellScript, powershellArguments, "", CallbackMockup)
	if err != nil {
		return result, err
	}

	return result, err
}

func DeleteSharedMailbox(appid string, ExchangeObjectId string) (result *EmptyResult, err error) {
	powershellScript := "scripts/sharedmailboxes/remove.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s`, ExchangeObjectId)
	result, err = RunExchange[EmptyResult](appid, powershellScript, powershellArguments, "", CallbackMockup)
	if err != nil {
		return result, err
	}

	return result, err
}

func UpdateSharedMailbox(appid string, ExchangeObjectId string, DisplayName string) (result *EmptyResult, err error) {
	powershellScript := "scripts/sharedmailboxes/update.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -DisplayName "%s"`, ExchangeObjectId, DisplayName)
	result, err = RunExchange[EmptyResult](appid, powershellScript, powershellArguments, "", CallbackMockup)
	if err != nil {
		return result, err
	}

	return result, err

}

func UpdateSharedMailboxPrimaryEmailAddress(appid string, ExchangeObjectId string, Email string) (result *EmptyResult, err error) {
	powershellScript := "scripts/sharedmailboxes/updateprimaryemail.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -Email "%s"`, ExchangeObjectId, Email)
	result, err = RunExchange[EmptyResult](appid, powershellScript, powershellArguments, "", CallbackMockup)
	if err != nil {
		return result, err
	}

	return result, err

}

func AddSharedMailboxMembers(appid string, ExchangeObjectId string, Members []string) (members *MembersResponse, err error) {
	powershellScript := "scripts/sharedmailboxes/addmembers.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -Members %s`, ExchangeObjectId, PwshArray(Members))
	members, err = RunExchange[MembersResponse](appid, powershellScript, powershellArguments, "", CallbackMockup)
	return members, err
}

func AddSharedMailboxReaders(appid string, ExchangeObjectId string, Readers []string) (members *MembersResponse, err error) {
	powershellScript := "scripts/sharedmailboxes/addreaders.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -Readers %s`, ExchangeObjectId, PwshArray(Readers))
	members, err = RunExchange[MembersResponse](appid, powershellScript, powershellArguments, "", CallbackMockup)
	return members, err
}

func SetSharedMailboxOwners(appid string, ExchangeObjectId string, Owners []string) (res *OwnersResponse, err error) {
	powershellScript := "scripts/sharedmailboxes/setowners.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -Owners %s`, ExchangeObjectId, PwshArray(Owners))
	res, err = RunExchange[OwnersResponse](appid, powershellScript, powershellArguments, "", CallbackMockup)
	return res, err
}

func RemoveSharedMailboxMembers(appid string, ExchangeObjectId string, Members []string) (members *MembersResponse, err error) {
	powershellScript := "scripts/sharedmailboxes/removemembers.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -Members %s`, ExchangeObjectId, PwshArray(Members))
	members, err = RunExchange[MembersResponse](appid, powershellScript, powershellArguments, "", CallbackMockup)
	return members, err
}

func RemoveSharedMailboxReaders(appid string, ExchangeObjectId string, Readers []string) (members *MembersResponse, err error) {
	powershellScript := "scripts/sharedmailboxes/removereaders.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -Readers %s`, ExchangeObjectId, PwshArray(Readers))
	members, err = RunExchange[MembersResponse](appid, powershellScript, powershellArguments, "", CallbackMockup)
	return members, err
}

func SetDistributionListMembers(appid string, upn string, requestForOnlyThisMemberships []string, withinThisDistributionListExchangeObjectIds []string) (*SyncResponse, error) {
	powershellScript := "scripts/distributionlists/setmemberships.ps1"
	powershellArguments := fmt.Sprintf(` -UPN %s -Memberships %s  -DistributionLists %s`, upn, PwshArray(requestForOnlyThisMemberships), PwshArray(withinThisDistributionListExchangeObjectIds))
	response, err := RunExchange[SyncResponse](appid, powershellScript, powershellArguments, "", CallbackMockup)
	return response, err

}
