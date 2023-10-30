package channels

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/model"
	"go.mongodb.org/mongo-driver/bson"
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

func GetScriptUpdateMembers(segment Segment, members []string) (string, error) {

	log.Println("Processing Mailgroup Segment", segment.Name)
	script := ""
	for _, value := range segment.Values {
		memberString := strings.Join(members, `","`)
		//members = members[0 : len(members)-1]
		script += fmt.Sprintf(`
Update-DistributionGroupMember  -Identity "zc-dl-%s" -Members "%s" -Confirm:$false	
	`, value.KeyHash, memberString)
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

func processMailGroupSegments(segments []Segment) error {

	log.Println("Processing Mailgroup Segments")

	for _, segment := range segments {
		script, err := GetScriptProcessMailGroupSegment(segment)
		if err != nil {
			return err
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
func processSMTPS(segments []Segment) (map[string]string, error) {

	log.Println("Processing SMTP's")

	smtps := getUniqueSMTPs(segments)

	// script := GetScriptCreateMailContacts(smtps)
	// data, err := model.ExecutePowerShellScript("get_guids", "exchange", script, "")

	script := GetGuidsForSMTPs(smtps)

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

/*
*

Updates the members of all mailgroups with the members from the segments.

This is done by mapping the SMTP address to the GUID of the user, and then use the GUID as a reference when
running the

Update-DistributionGroupMember
*/
func updateMembers(segments []Segment, smtp2guidMap map[string]string) error {

	log.Println("Updating members")

	for _, segment := range segments {

		members := make([]string, 0)
		for _, value := range segment.Values {
			for _, smtp := range value.Values {
				if guid, ok := smtp2guidMap[strings.ToLower(smtp)]; ok {
					members = append(members, guid)
				} else {
					log.Println("Could not find guid for", smtp)
				}
			}
		}
		script, err := GetScriptUpdateMembers(segment, members)
		if err != nil {
			return err
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
func SyncDistributionGroups() (err error) {
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
	err = processMailGroupSegments(segmentdata.Segments)
	if err != nil {
		return err
	}
	smtp2guidMap, err := processSMTPS(segmentdata.Segments)
	if err != nil {
		return err
	}
	err = updateMembers(segmentdata.Segments, smtp2guidMap)
	if err != nil {
		return err
	}

	//ExportIndex(segmentdata.Segments)
	// Publish new Group Finder map
	log.Println("Done")

	return nil
}

func Sync() error {
	return SyncDistributionGroups()
}
