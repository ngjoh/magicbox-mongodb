package scaffold

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetTsField(field Field) (string, string) {
	switch field.Type {
	case "Text":
		return fmt.Sprintf(`item.fields.%s ? item.fields.%s : ""`, field.Name, field.Name),
			fmt.Sprintf(`%s : z.string()`, field.Name)
	case "Note":
		return fmt.Sprintf(`item.fields.%s ? item.fields.%s : ""`, field.Name, field.Name),
			fmt.Sprintf(`%s : z.string()`, field.Name)
	case "Number":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.number()`, field.Name)
	case "Boolean":
		return fmt.Sprintf(`item.fields.%s ? true : false`, field.Name),
			fmt.Sprintf(`%s : z.boolean()`, field.Name)
	case "DateTime":
		return fmt.Sprintf(`new Date(item.fields.%s)`, field.Name),
			fmt.Sprintf(`%s : z.date()`, field.Name)
	case "User":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name)
	case "Lookup":
		return fmt.Sprintf(`mapLookup(item.fields.%s)`, field.Name),
			fmt.Sprintf(`%s : z.object({
				LookupId:z.number(),
				LookupValue:z.string()
			  }).nullable()`, field.Name)
	case "LookupMulti":
		return fmt.Sprintf(`mapLookupMulti(item.fields.%s)`, field.Name),
			fmt.Sprintf(`%s : z.object({
				LookupId:z.number(),
				LookupValue:z.string()
			  }).array().nullable()`, field.Name)
	case "URL":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name)
	case "Choice":

		return fmt.Sprintf(`item.fields.%s ?? ""`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name)
	case "MultiChoice":
		return fmt.Sprintf(`item.fields.%s ?? []`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name)
	case "Calculated":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name)
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

	return "", ""
}

func OutTS(list ListInstance) string {

	itemsMap := `
	export function map(item:any) {
	
	
	return {
		Id : item.id,
	Title : item.fields.Title,
	CreatedBy : item.createdBy.user.email,
	Created :new Date(item.createdDateTime),
	ModifiedBy : item.lastModifiedBy.user.email,
	Modified : new Date(item.lastModifiedDateTime),	
		`

	zodMaps := `
	export const schema = z.object({
		CreatedBy : z.string(),
		Created: z.date(),
		ModifiedBy : z.string(),
		Modified: z.date(),
		Id: z.string(),
		Title: z.string(),
		
		`

	for _, field := range list.Fields.Field {
		fieldMap, zodMap := GetTsField(field)

		if fieldMap != "" {
			itemsMap += fmt.Sprintf(`%s: %s,
			`, field.Name, fieldMap)

		}
		if zodMap != "" {
			zodMaps += fmt.Sprintf(`%s,
			`, zodMap)

		}

	}
	itemsMap += "}}"
	zodMaps += "})"
	sharepointMap := ""

	sharepointMap += fmt.Sprintf(`export namespace %s {
		export const listName = "%s"
		`, strings.ReplaceAll(list.Title, " ", ""), list.Title)
	sharepointMap += itemsMap

	sharepointMap += zodMaps
	sharepointMap += "}"
	return sharepointMap

}

func Pnp2Ts(filename string) string {

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
	sharepointMap := `import z from "zod"
	// Generated by pnp2ts - do not edit


	function mapLookup(item:any) {
		if (!item) return null
		return {
			LookupId: item.LookupId,
			LookupValue: item.LookupValue
		}
	}

	function mapLookupMulti(items:any[]) {
		if (!items) return []
		return items.map(item =>
		{
			return {LookupId: item.LookupId,
				LookupValue: item.LookupValue}
		})
	}

	`
	for _, list := range template.Templates.ProvisioningTemplate.Lists.ListInstance {
		sharepointMap += OutGo(list)
	}

	return sharepointMap

}
