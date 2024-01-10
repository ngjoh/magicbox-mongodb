package kitchen

import "testing"

func TestCreateWorkspaceFile(t *testing.T) {
	err := CreateWorkspaceFile()
	if err != nil {
		t.Error(err)
	}

}

func TestCreateKitchen(t *testing.T) {
	err := CreateKitchen("nexi-sharepoint")
	if err != nil {
		t.Error(err)
	}

}
