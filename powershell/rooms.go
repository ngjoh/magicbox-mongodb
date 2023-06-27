package powershell

import (
	"fmt"
)

type NewRoomResult struct {
	MailAddress string `json:"MailAddress"`
}

func CreateRoom(Name string, Capacity int) (result *NewRoomResult, err error) {
	powershellScript := "scripts/rooms/create-room.ps1"
	powershellArguments := fmt.Sprintf(` -Name "%s" -Capacity %d `, Name, Capacity)
	result, err = RunExchange[NewRoomResult](powershellScript, powershellArguments)
	if err != nil {
		return result, err
	}

	return result, err
}

func RemoveRoom(Email string) (result *EmptyResult, err error) {
	powershellScript := "scripts/rooms/remove-room.ps1"
	powershellArguments := fmt.Sprintf(` -Mail "%s" `, Email)
	result, err = RunExchange[EmptyResult](powershellScript, powershellArguments)
	if err != nil {
		return result, err
	}

	return result, err
}
