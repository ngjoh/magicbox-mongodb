package kitchen

import (
	"testing"
)

func TestCreateWorkspaceFile(t *testing.T) {
	err := CreateWorkspaceFile()
	if err != nil {
		t.Error(err)
	}

}

func TestCreateKitchen(t *testing.T) {
	err := CreateKitchen("teams-admin")
	if err != nil {
		t.Error(err)
	}

}

func TestCreateKitchen2(t *testing.T) {
	err := CreateKitchen("365admin-v1")
	if err != nil {
		t.Error(err)
	}

}
func TestBuildKitchen(t *testing.T) {
	err := BuildKitchen("azure-users")
	if err != nil {
		t.Error(err)
	}

}

func TestAddTemplate(t *testing.T) {
	err := AddTemplatesToKitchen("azure-roles")
	if err != nil {
		t.Error(err)
	}
}
func TestGetKitchenStatus(t *testing.T) {
	status, err := GetKitchenStatus("meeting-infrastructure")
	if err != nil {
		t.Error(err)
	}
	t.Log(status)
}

func TestMakeRelease(t *testing.T) {

	err := MakeRelease("meeting-infrastructure", "v0.0.2")
	if err != nil {
		t.Error(err)
	}

}
