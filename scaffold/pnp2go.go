package scaffold

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func json(name string) string {
	return fmt.Sprintf(`"json:%s"`, name)
}
func getListName(list string, replaceSpaceWith string) string {
	return strings.Replace(strings.Replace(strings.Replace(list, "{listid:", "", -1), "}", "", -1), " ", replaceSpaceWith, -1)
}

// Return input and output model
func GetGoField(field Field) (string, string, string, string) {
	if field.Name == "_ModernAudienceAadObjectIds" {
		return "", "", "", ""
	}
	switch field.Type {
	case "Text":
		return Pair(cases.Title(language.English).String(field.Name), "string", json(field.Name)), Pair(field.Name, "string", json(field.Name)), field.Name, ""

	case "Note":
		return Pair(cases.Title(language.English).String(field.Name), "string", json(field.Name)), Pair(field.Name, "float64", json(field.Name)), field.Name, ""
	case "Number":
		return Pair(cases.Title(language.English).String(field.Name), "float64", json(field.Name)), Pair(field.Name, "string", json(field.Name)), field.Name, ""
	case "Int":
		return Pair(cases.Title(language.English).String(field.Name), "int", json(field.Name)), Pair(field.Name, "int", json(field.Name)), field.Name, ""
	case "Boolean":
		return Pair(cases.Title(language.English).String(field.Name), "bool", json(field.Name)), Pair(field.Name, "bool", json(field.Name)), field.Name, ""
	case "DateTime":
		return Pair(cases.Title(language.English).String(field.Name), "time.Time", json(field.Name)), Pair(field.Name, "time.Time", json(field.Name)), field.Name, ""
	case "User":
		return "", "", "", ""

	// case "Lookup":  // not supported (yet)
	// 	return Pair(field.Name, "shared.LookupReference", json(field.Name)), Pair(field.Name, getListName(field.List, ""), json(field.Name)), field.Name, getListName(field.List, " ")

	case "LookupMulti":
		return Pair(cases.Title(language.English).String(field.Name), "[]shared.LookupReference", json(field.Name)), Pair(field.Name, getListName(field.List, ""), json(field.Name)), field.Name, getListName(field.List, " ")

	case "URL":
		return "", "", "", ""

	case "Choice":
		return Pair(cases.Title(language.English).String(field.Name), "string", json(field.Name)), "", field.Name, ""

	case "MultiChoice":
		return Pair(cases.Title(language.English).String(field.Name), "[]string", json(field.Name)), "", field.Name, ""

	case "Calculated":
		return Pair(cases.Title(language.English).String(field.Name), "string", json(field.Name)), "", field.Name, ""

	// case "Attachments":
	// 	return "string"
	// case "Guid":
	// 	return "string"
	// case "Integer":
	// 	return "number"
	// case "Counter":

	// 	return "number"
	// case "Currency":

	// 	return "number"
	default:
		// log.Println("Unknown type", field.Type)
	}

	return "", "", "", ""
}
func LookupMultiStruct() string {
	return fmt.Sprintf(`struct {
		Results []int %s
	} `, json("results"))
}
func Pair(fieldName string, typeName string, mapping string) string {
	return fmt.Sprintf("%s %s %s%s%s", fieldName, typeName, "`", mapping, "`")
}
func getPublicName(fieldName string) string {
	return strings.ReplaceAll(cases.Title(language.English).String(fieldName), " ", "")
}
func CommonField(name string, typeName string) string {
	field := &Field{Name: name, Type: typeName}
	v, _, _, _ := GetGoField(*field)
	return fmt.Sprintf("%s %s", v, `
	`)
}
func OutGo(list ListInstance) string {
	dependencies := map[string]string{}

	itemsMap := ""

	itemsMap += CommonField("Title", "Text")
	// itemsMap += CommonField("Created", "DateTime")
	// itemsMap += CommonField("EditorId", "Int")
	// itemsMap += CommonField("GUID", "Text")
	// itemsMap += CommonField("ID", "Int")
	// itemsMap += CommonField("Modified", "DateTime")
	// itemsMap += CommonField("AuthorId", "Int")

	//  {
	// 	mgm.DefaultModel %s
	// 	Id : item.id,
	// Title : item.fields.Title,
	// CreatedBy : item.createdBy.user.email,
	// Created :new Date(item.createdDateTime),
	// ModifiedBy : item.lastModifiedBy.user.email,
	// Modified : new Date(item.lastModifiedDateTime),
	// 	`,"`bson:\",inline\"`")

	zodMaps := ""
	zodMaps += CommonField("Title", "Text")
	zodMaps += CommonField("Created", "DateTime")
	zodMaps += CommonField("EditorId", "Int")
	zodMaps += CommonField("GUID", "Text")
	zodMaps += CommonField("ID", "Int")
	zodMaps += CommonField("Modified", "DateTime")
	zodMaps += CommonField("AuthorId", "Int")
	fieldArray := []string{}

	for _, field := range list.Fields.Field {
		fieldMap, zodMap, fieldNameToSelect, dependency := GetGoField(field)
		if fieldNameToSelect != "" {
			fieldArray = append(fieldArray, fmt.Sprintf(`"%s"`, fieldNameToSelect))
		}
		if dependency != "" {
			d := fmt.Sprintf(`"%s"`, dependency)
			_, ok := dependencies[d]
			if !ok {
				dependencies[d] = d
			}
		}
		if fieldMap != "" {
			itemsMap += fmt.Sprintf(`%s
			`, fieldMap)

		}
		if zodMap != "" {
			zodMaps += fmt.Sprintf(`%s
			`, zodMap)

		}

	}
	itemsMap += ""
	zodMaps += ""
	sharepointMap := ""

	publicName := fmt.Sprintf("SP_%s", getPublicName(list.Title))
	sharepointMap += fmt.Sprintf(`
		type  %s struct {
		`, publicName)
	sharepointMap += itemsMap

	sharepointMap += "}"

	publicFields := fmt.Sprintf("SP_%s_Fields", getPublicName(list.Title))
	sharepointMap += fmt.Sprintf(`
		var  %s = []string{`, publicFields)
	sharepointMap += strings.Join(fieldArray, ",")

	sharepointMap += "}"

	publicFields = fmt.Sprintf("SP_%s_Dependencies", getPublicName(list.Title))
	sharepointMap += fmt.Sprintf(`
		var  %s = []string{`, publicFields)

	keys := []string{}

	for key := range dependencies {
		keys = append(keys, key)
	}
	sharepointMap += strings.Join(keys, ",")

	sharepointMap += "}"

	dbName := getPublicName(list.Title)
	sharepointMap += fmt.Sprintf(`
		type  %s struct {
		`, dbName)
	sharepointMap += zodMaps

	sharepointMap += "}"
	return sharepointMap

}
func Pnp2Go(filename string) string {

	// Open our xmlFile
	xmlFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Successfully Opened template.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var template Provisioning
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &template)

	// we iterate through every list within our template
	sharepointMap := `package nexiintra_home
	// Generated by pnp2go - do not edit
	import (
		"time"
		"github.com/koksmat-com/koksmat/shared"
		
	)
	
	`

	for _, list := range template.Templates.ProvisioningTemplate.Lists.ListInstance {
		sharepointMap += OutGo(list)
	}

	return sharepointMap

}
