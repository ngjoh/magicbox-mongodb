package queue

import (
	"encoding/json"
	"fmt"

	"github.com/koksmat-com/koksmat/worker"
)

func PowerShellRequestResponse[R any](appId string, scriptName string, input string) (result *R, err error) {
	dataOut := new(R)
	response, err := worker.PowerShellExecute(appId, scriptName, input)
	if err != nil {
		return result, err
	}
	if response.HasError {
		return result, fmt.Errorf(response.Output)
	}
	err = json.Unmarshal([]byte(response.Output), &dataOut)
	result = *&dataOut
	return result, err
}
