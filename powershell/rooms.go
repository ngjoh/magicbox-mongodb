package powershell

import (
	"fmt"
	"strings"
)

type NewRoomResult struct {
	MailAddress string `json:"MailAddress"`
}

func CreateRoom(Name string, Capacity int) (result *NewRoomResult, err error) {
	powershellScript := "scripts/rooms/create-room.ps1"
	powershellArguments := fmt.Sprintf(` -Name "%s" -Capacity %d `, Name, Capacity)
	result, err = RunExchange[NewRoomResult]("koksmat", powershellScript, powershellArguments, "", CallbackMockup)
	if err != nil {
		return result, err
	}

	return result, err
}

func RemoveRoom(Email string) (result *EmptyResult, err error) {
	powershellScript := "scripts/rooms/remove-room.ps1"
	powershellArguments := fmt.Sprintf(` -Mail "%s" `, Email)
	result, err = RunExchange[EmptyResult]("koksmat", powershellScript, powershellArguments, "", CallbackMockup)
	if err != nil {
		return result, err
	}

	return result, err
}

func BuildPolicyScriptArguments(Email string, ListOfAllowedEmails []string) (string, string) {
	sharedArguments := fmt.Sprintf(` -Mail "%s" -BookingWindowInDays 601`, Email)

	if len(ListOfAllowedEmails) == 0 {
		return "scripts/rooms/set-roompolicy-standard.ps1", fmt.Sprintf(`%s `, sharedArguments)
	} else {
		return "scripts/rooms/set-roompolicy-restricted.ps1", fmt.Sprintf(`%s -MailTip "This room has restrictions on who can book it" -AllowedBooker %s `, sharedArguments, PwshArray(ListOfAllowedEmails))
	}

}

func SetPolicy(Email string, ListOfAllowedEmails string) (result *EmptyResult, err error) {

	powershellScript, powershellArguments := BuildPolicyScriptArguments(Email, strings.Split(ListOfAllowedEmails, ","))
	return RunExchange[EmptyResult]("koksmat", powershellScript, powershellArguments, "", CallbackMockup)

}
