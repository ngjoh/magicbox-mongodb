package powershell

import (
	"fmt"
)

type NewSharedMailboxResult struct {
	Name               string `json:"Name"`
	DisplayName        string `json:"DisplayName"`
	ExchangeObjectId   string `json:"ExchangeObjectId"`
	PrimarySmtpAddress string `json:"PrimarySmtpAddress"`
}
type EmptyResult struct {
}

func CreateSharedMailbox(Name string, DisplayName string, Alias string, Owners []string, Members []string, Readers []string) (result *NewSharedMailboxResult, err error) {
	powershellScript := "scripts/sharedmailboxes/create.ps1"
	powershellArguments := fmt.Sprintf(` -Name "test5-%s" -DisplayName "Test5 %s"  -Alias "test5-%s" -Members "%s" -Readers "%s" -Owner="%s"`, Name, DisplayName, Alias, Members, Readers, Owners)
	result, err = Run[NewSharedMailboxResult](powershellScript, powershellArguments)
	if err != nil {
		return result, err
	}

	return result, err
}

func DeleteSharedMailbox(ExchangeObjectId string) (result *EmptyResult, err error) {
	powershellScript := "scripts/sharedmailboxes/remove.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s`, ExchangeObjectId)
	result, err = Run[EmptyResult](powershellScript, powershellArguments)
	if err != nil {
		return result, err
	}

	return result, err
}

func UpdateSharedMailbox(ExchangeObjectId string, DisplayName string) (result *EmptyResult, err error) {
	powershellScript := "scripts/sharedmailboxes/update.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -DisplayName "%s"`, ExchangeObjectId, DisplayName)
	result, err = Run[EmptyResult](powershellScript, powershellArguments)
	if err != nil {
		return result, err
	}

	return result, err

}

func AddSharedMailboxMembers(ExchangeObjectId string, Members []string) (err error) {
	powershellScript := "scripts/sharedmailboxes/addmembers.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -Members "%s"`, ExchangeObjectId, Members)
	_, err = Run[EmptyResult](powershellScript, powershellArguments)
	return err
}

func AddSharedMailboxReaders(ExchangeObjectId string, Members []string) (err error) {
	powershellScript := "scripts/sharedmailboxes/addreaders.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -Readers "%s"`, ExchangeObjectId, Members)
	_, err = Run[EmptyResult](powershellScript, powershellArguments)
	return err
}

func AddSharedMailboxOwners(ExchangeObjectId string, Owners []string) (err error) {
	powershellScript := "scripts/sharedmailboxes/addowners.ps1"
	powershellArguments := fmt.Sprintf(` -ExchangeObjectId %s -Owners "%s"`, ExchangeObjectId, Owners)
	_, err = Run[EmptyResult](powershellScript, powershellArguments)
	return err
}
