package channels

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Value struct {
	Key     string   `json:"key"`
	KeyHash string   `json:"keyHash"`
	Values  []string `json:"values"`
}
type Segment struct {
	Name   string  `json:"name"`
	Values []Value `json:"values"`
}
type MailgroupsSegmentdata struct {
	mgm.DefaultModel `bson:",inline"`
	Version          string    `json:"version"`
	Columns          []string  `json:"columns"`
	Segments         []Segment `json:"segments"`
}

func GetScriptProcessMailGroupSegment(segment Segment) (string, error) {

	log.Println("Processing Mailgroup Segment", segment.Name)
	script := ""
	for _, value := range segment.Values {
		script += fmt.Sprintf(`
$upn = "niels.johansen@nexigroup.com"

$ErrorActionPreference = "SilentlyContinue"
Write-Out Checking "%s" 
$dl = Get-DistributionGroup "zc-dl-%s"
if ($dl -eq $null){
	$ErrorActionPreference = "Continue"
	Write-Out New-DistributionGroup -Name "zc-dl-%s" -DisplayName "%s" -ManagedBy ${upn} 
	New-DistributionGroup -Name "zc-dl-%s" -DisplayName "%s" -ManagedBy ${upn} 
	$ErrorActionPreference = "SilentlyContinue"
}	
	
	`, value.Key, value.KeyHash, value.KeyHash, value.Key, value.KeyHash, value.Key)
	}
	return script, nil
}

func processMailGroupSegments(segments []Segment) {

	log.Println("Processing Mailgroup Segments")

	for _, segment := range segments {
		script, err := GetScriptProcessMailGroupSegment(segment)
		if err != nil {
			log.Println(err)
		}

		_, err = model.ExecutePowerShellScript("tester", "exchange", script, "")

		if err != nil {
			log.Fatalf("processMailGroupSegment() error = %v", err)
		}
	}

}

func CreateNewDistributionGroups() (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	results, err := mgm.Coll(&MailgroupsSegmentdata{}).Find(context.TODO(), bson.M{})
	if err != nil {

		return err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var segmentdata MailgroupsSegmentdata
		if err = results.Decode(&segmentdata); err != nil {
			return err
		}
		log.Println("Loaded ", segmentdata.Version, " Mailgroup Segments")
		processMailGroupSegments(segmentdata.Segments)

	}
	return nil
}
