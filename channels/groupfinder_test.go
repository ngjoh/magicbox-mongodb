package channels

import (
	"testing"

	"github.com/koksmat-com/koksmat/model"
)

func TestSyncDistributionGroups(t *testing.T) {

	if err := SyncDistributionGroups(true); err != nil {
		t.Errorf("CreateNewDistributionGroups() error = %v", err)
	}

}

func TestGetAllZCMailgroups(t *testing.T) {

	if _, err := GetAllZCMailgroups(); err != nil {
		t.Errorf("GetAllZCMailgroups() error = %v", err)
	}

}
func TestAttachSMTPmap(t *testing.T) {

	if err := attachSMTPmap(); err != nil {
		t.Errorf("attachSMTPmap() error = %v", err)
	}

}
func CallbackMockup(workingDirectory string) {}
func TestProcessMailGroupSegment(t *testing.T) {
	segment := &Segment{
		Name: "Test",
		Values: []Value{
			{
				Key:     "All employees Centrum Rozliczeń El. [Nets]",
				KeyHash: "company-31de442c0e5c09d941d23300f0819e8536bb5f63",
				Values:  []string{"MARTYNA.MROCZKA@PEP.PL", "KAMILA.KALANDYK@PEP.PL"},
			}},
	}

	powershellScript, err := GetScriptProcessMailGroupSegment(*segment)
	if err != nil {
		t.Errorf("processMailGroupSegment() error = %v", err)
	}

	_, err = model.ExecutePowerShellScript("tester", "exchange", powershellScript, "")

	if err != nil {
		t.Errorf("processMailGroupSegment() error = %v", err)
	}
}

// func TestProcessMailgroupsSegments(t *testing.T) {
// 	processMailGroupSegments()
// }
