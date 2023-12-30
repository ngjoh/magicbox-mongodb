package channels

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/config"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/model"
	"github.com/koksmat-com/koksmat/powershell"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/slices"
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
	Version          string        `json:"version"`
	Columns          []string      `json:"columns"`
	Segments         []Segment     `json:"segments"`
	SmtpMap          SmtpToGuidMap `json:"smtpMap"`
}

func GetAllZCMailgroups() ([]string, error) {
	guidMap := make([]string, 0)
	script := `
	$result = Get-DistributionGroup "zc-dl" -ResultSize 10 | select -ExpandProperty PrimarySmtpAddress

	ConvertTo-Json  -InputObject $result -Depth 10
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
	`

	data, err := model.ExecutePowerShellScript("GetAllZCMailgroups", "exchange", script, "")

	err = json.Unmarshal([]byte(data), &guidMap)
	if err != nil {
		return nil, err
	}
	return guidMap, nil
}

func GetScriptProcessMailGroupSegment(segment Segment) (string, error) {

	log.Println("Processing Mailgroup Segment", segment.Name)
	script := ""
	for _, value := range segment.Values {
		script += fmt.Sprintf(`
$upn = "niels.johansen@nexigroup.com"

$ErrorActionPreference = "SilentlyContinue"
Write-Host Checking "%s" 
$dl = Get-DistributionGroup "zc-dl-%s"
if ($dl -eq $null){
	$ErrorActionPreference = "Continue"
	Write-Host New-DistributionGroup -Name "zc-dl-%s" -DisplayName "%s" -ManagedBy ${upn} 
	New-DistributionGroup -Name "zc-dl-%s" -DisplayName "%s" -ManagedBy ${upn} 
	$ErrorActionPreference = "SilentlyContinue"
}	
	
	`, value.Key, value.KeyHash, value.KeyHash, value.Key, value.KeyHash, value.Key)
	}
	return script, nil
}
func GetScriptCreateMailContacts(smtps []string) string {

	script := ""
	for _, smtp := range smtps {
		script += fmt.Sprintf(`

$ErrorActionPreference = "SilentlyContinue"
$smtp =  "%s" 
Write-Host Checking $smtp
$recipient = Get-Recipient $smtp
if ($recipient -eq $null){
	write-host "Not found $smtp"
	if ($smtp.ToLower().IndexOf("nexigroup.com") -eq -1){
		$ErrorActionPreference = "Continue"
		Write-Host New-MailContact -ExternalEmailAddress $smtp  -Name $smtp
		New-MailContact -ExternalEmailAddress $smtp  -Name $smtp | fl
		$ErrorActionPreference = "SilentlyContinue"
	}else{
		Write-Host "Not creating $smtp"
	}

}	
	
$ErrorActionPreference = "Continue"
`, smtp)
	}
	return script
}

func GetGuidsForSMTPs(smtps []string) string {

	script := `
$result = @()	
	`
	for ix, smtp := range smtps {

		script += fmt.Sprintf(`
$ErrorActionPreference = "SilentlyContinue"
$smtp =  "%s" 

Write-Host Checking %d of %d : $smtp
$recipient = Get-Recipient $smtp -ResultSize 1
if ($null -ne $recipient){
		$result += "$smtp=$($recipient.Guid)"
}	


			
`, smtp, ix, len(smtps))
	}
	script += `
ConvertTo-Json  -InputObject $result -Depth 10
| Out-File -FilePath $PSScriptRoot/output.json -Encoding:utf8NoBOM
		
	`
	return script
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func GetScriptUpdateMembers(dryrun bool, segment Segment, smtp2guidMap map[string]string) (string, error) {

	log.Println("Processing Mailgroup Segment", segment.Name)
	script := ""
	for ix, value := range segment.Values {
		members := []string{}
		for _, smtp := range value.Values {
			if dryrun {
				members = append(members, smtp) // assume match
				continue
			}
			if guid, ok := smtp2guidMap[strings.ToLower(smtp)]; ok {
				members = append(members, guid)
			} else {
				log.Println("Could not find guid for", smtp)
			}
		}
		memberString := strings.Join(removeDuplicateStr(members), `","`)
		log.Println(fmt.Sprintf(`%d of %d - %d members i "zc-dl-%s" `, ix+1, len(segment.Values), len(members), value.Key))
		//members = members[0 : len(members)-1]
		script += fmt.Sprintf(`
$ErrorActionPreference = "Continue"
Write-Host " %d of %d - %d members of %s  (zc-dl-%s)"
Update-DistributionGroupMember  -Identity "zc-dl-%s" -Members "%s" -Confirm:$false	
	`, ix+1, len(segment.Values), len(members), value.Key, value.KeyHash, value.KeyHash, memberString)
	}
	return script, nil
}
func getUniqueSMTPs(segments []Segment) []string {
	smtps := make([]string, 0)
	for _, segment := range segments {
		for _, value := range segment.Values {
			for _, smtp := range value.Values {

				if !slices.Contains(smtps, smtp) {
					smtps = append(smtps, smtp)
				}

			}
		}
	}
	return smtps
}

func processMailGroupSegments(dryrun bool, workingDirectory string, segments []Segment) error {

	log.Println("Processing Mailgroup Segments")
	log.Println("Dryrun, skipping execution of script")
	for _, segment := range segments {
		script, err := GetScriptProcessMailGroupSegment(segment)
		log.Println("Processing Mailgroup Segment", segment.Name)
		if err != nil {
			return err
		}
		err = os.WriteFile(path.Join(workingDirectory, fmt.Sprintf("check-groups-%s.ps1", url.QueryEscape(segment.Name))), []byte(string(script)), 0644)
		if dryrun {

			continue
		}

		_, err = model.ExecutePowerShellScript("create_missing_groups", "exchange", script, "")

		if err != nil {
			return err
		}
	}
	return nil
}

type SmtpToGuidMap []string

/*
ProcessSMTPS processes an array of smtp addresses by generating
a powershell script that will query the Exchange server for the guid of each smtp address.

It returns a map of smtp address to guid. to that you can convert the smtp address to guid.

Converting a smtp address to a GUID is necessary when you want to add a user to a distribution group,
as there is no guarantee that the smtp address is unique across all types of objects in the Exchange server.

Notice that the map returned contains the smtp address in lowercase as the key, and the GUID as the value.
*/
func processSMTPS(dryrun bool, workingDirectory string, segments []Segment) (map[string]string, error) {

	log.Println("Processing SMTP's")

	smtps := getUniqueSMTPs(segments)

	// script := GetScriptCreateMailContacts(smtps)
	// data, err := model.ExecutePowerShellScript("get_guids", "exchange", script, "")

	script := GetGuidsForSMTPs(smtps)
	err := os.WriteFile(path.Join(workingDirectory, "process-smtp.ps1"), []byte(string(script)), 0644)
	if dryrun {
		log.Println("Dryrun, skipping execution of script")
		return nil, nil
	}

	data, err := model.ExecutePowerShellScript("get_guids", "exchange", script, "")
	if err != nil {
		return nil, err
	}
	guidMap := SmtpToGuidMap{}
	err = json.Unmarshal([]byte(data), &guidMap)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	for _, guid := range guidMap {
		pair := strings.Split(guid, "=")

		m[strings.ToLower(pair[0])] = pair[1]
	}

	log.Println("Got", len(guidMap), "guids from", len(smtps), "smtps")

	return m, nil

}
func buildSMTPSmap(guidMap SmtpToGuidMap) (map[string]string, error) {
	m := make(map[string]string)
	for _, guid := range guidMap {
		pair := strings.Split(guid, "=")

		m[strings.ToLower(pair[0])] = pair[1]
	}

	log.Println("Got", len(guidMap), "guids from", len(guidMap), "smtps")
	return m, nil

}
func attachSMTPmap() error {
	data, err := os.ReadFile("/Users/nielsgregersjohansen/code/koksmat/koksmat-cli/.koksmat/powershell/get_guids-d850fadb-8cb6-4353-9d4c-a84c20149809/output.json")
	var smtpMap SmtpToGuidMap
	json.Unmarshal(data, &smtpMap)
	ctx := context.TODO()

	// Connect to MongoDB
	client := db.Connect()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Open an aggregation cursor

	coll := client.Database(config.DatabaseName()).Collection("mailgroups_segmentdata")
	id, _ := primitive.ObjectIDFromHex("653a2ab5eed3db63704f1871")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"smtpMap", smtpMap}}}}

	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

/*
*

Updates the members of all mailgroups with the members from the segments.

This is done by mapping the SMTP address to the GUID of the user, and then use the GUID as a reference when
running the

Update-DistributionGroupMember
*/
func updateMembers(dryrun bool, workingDirectory string, segments []Segment, smtp2guidMap map[string]string) error {

	log.Println("Updating members")
	log.Println("Dryrun, skipping execution of script")
	for _, segment := range segments {
		log.Printf("Updating members for %s", segment.Name)

		script, err := GetScriptUpdateMembers(dryrun, segment, smtp2guidMap)
		if err != nil {
			return err
		}
		err = os.WriteFile(path.Join(workingDirectory, fmt.Sprintf("update-members-%s.ps1", url.QueryEscape(segment.Name))), []byte(string(script)), 0644)
		if dryrun {

			continue
		}
		_, err = model.ExecutePowerShellScript("update_members", "exchange", script, "")

		if err != nil {
			return err
		}
	}
	return nil
}

type GroupFinderIndex struct {
	HasError bool `json:"hasError"`
	Data     struct {
		Sheets  []any `json:"sheets"`
		Columns []any `json:"columns"`
		Results struct {
			OnSheetLoaded struct {
				Version  string    `json:"version"`
				Columns  []string  `json:"columns"`
				Segments []Segment `json:"segments"`
			} `json:"onSheetLoaded"`
		} `json:"results"`
	} `json:"data"`
}

/**
Iterate over all segmnets, and create a new segment that only contains the name (Key) and alias (Keyhash) of each segment.

*/

func getIndexAllSegments(segments []Segment) []Segment {
	segmentsForIndex := make([]Segment, 0)
	for _, segment := range segments {
		segmentForIndex := Segment{
			Name: segment.Name,
		}
		for _, value := range segment.Values {
			valueForIndex := Value{
				Key:     value.Key,
				KeyHash: value.KeyHash,
			}
			segmentForIndex.Values = append(segmentForIndex.Values, valueForIndex)

		}
		segmentsForIndex = append(segmentsForIndex, segmentForIndex)
	}

	return segmentsForIndex
}

/*
*
 */

func ExportIndex(segments []Segment) error {
	index := &GroupFinderIndex{}
	index.HasError = false
	index.Data.Sheets = []any{}
	index.Data.Results.OnSheetLoaded.Version = "1"
	index.Data.Results.OnSheetLoaded.Segments = getIndexAllSegments(segments)

	return nil
}

/*
SyncDistributionGroups processes a list of mailgroup segments and creates the mailgroups and mailcontacts in Exchange
for those missing, and finally updates the members of all known mailgroups with the members from the segments.
*/
func SyncDistributionGroups(dryrun bool) (err error) {
	workingDirectory := powershell.PwshCwd("sync-distribution-groups")
	result := mgm.Coll(&MailgroupsSegmentdata{}).FindOne(context.TODO(), bson.D{{"processed", "none"}})
	if result.Err() != nil {
		log.Println("No document found")
		return nil
	}

	var segmentdata MailgroupsSegmentdata
	if err = result.Decode(&segmentdata); err != nil {
		return err
	}
	log.Println("Loaded ", segmentdata.Version, " Mailgroup Segments")
	err = processMailGroupSegments(dryrun, workingDirectory, segmentdata.Segments)
	if err != nil {
		return err
	}
	smtp2guidMap, err := processSMTPS(dryrun, workingDirectory, segmentdata.Segments)
	if err != nil {
		return err
	}

	//smtp2guidMap, _ := buildSMTPSmap(segmentdata.SmtpMap)

	err = updateMembers(dryrun, workingDirectory, segmentdata.Segments, smtp2guidMap)
	if err != nil {
		log.Println("Error updating members", err)
		return err
	}

	//ExportIndex(segmentdata.Segments)
	// Publish new Group Finder map

	return nil
}

func Sync(dryrun bool) error {
	log.Println("Syncing Distribution Groups", dryrun)

	return SyncDistributionGroups(dryrun)
}
