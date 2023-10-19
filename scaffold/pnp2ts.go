package scaffold

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func createProject(appId string) string {

	//dir := ".koksmat/sharepoint"
	dir := "/Users/nielsgregersjohansen/code/koksmat/ui/apps/www/app/sharepoint/models"
	os.MkdirAll(dir, os.ModePerm)
	dir = path.Join(dir, appId)
	os.MkdirAll(dir, os.ModePerm)

	return dir
}
func createSubdir(projDir string, subDir string) string {

	dir := path.Join(projDir, subDir)
	os.MkdirAll(dir, os.ModePerm)

	return dir
}
func writeFile(fileDir string, filename string, content string) {

	file, err := os.Create(path.Join(fileDir, filename))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString(content)
}
func GetTsField(field Field) (string, string, string) {
	switch field.Type {
	case "Text":
		return fmt.Sprintf(`item.fields.%s ? item.fields.%s : ""`, field.Name, field.Name),
			fmt.Sprintf(`%s : z.string()`, field.Name), fmt.Sprintf(`<FormField
			control={form.control}
			name="%s"
			render={({ field }) => (
				<FormItem>
					<FormLabel>%s</FormLabel>
					<FormControl>
						<Input placeholder="" {...field} />
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}
		/>`, field.Name, field.DisplayName, field.Description)
	case "Note":
		return fmt.Sprintf(`item.fields.%s ? item.fields.%s : ""`, field.Name, field.Name),
			fmt.Sprintf(`%s : z.string()`, field.Name), ""
	case "Number":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.number()`, field.Name), ""
	case "Boolean":
		return fmt.Sprintf(`item.fields.%s ? true : false`, field.Name),
			fmt.Sprintf(`%s : z.boolean()`, field.Name), ""
	case "DateTime":
		return fmt.Sprintf(`new Date(item.fields.%s)`, field.Name),
			fmt.Sprintf(`%s : z.date()`, field.Name), ""
	case "User":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), ""
	case "Lookup":
		return fmt.Sprintf(`mapLookup(%s,item.fields.%sLookupId)`, strings.ReplaceAll(strings.ReplaceAll(field.List, "{listid:", "\""), "}", "\""), field.Name),
			fmt.Sprintf(`%s : z.object({
				LookupId:z.number(),
				LookupValue:z.string()
			  }).nullable()`, field.Name), ""
	case "LookupMulti":
		return fmt.Sprintf(`mapLookupMulti(%s,item.fields.%sLookupId)`, strings.ReplaceAll(strings.ReplaceAll(field.List, "{listid:", "\""), "}", "\""), field.Name),
			fmt.Sprintf(`%s : z.object({
				LookupId:z.number(),
				LookupValue:z.string()
			  }).array().nullable()`, field.Name), ""
	case "URL":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), ""
	case "Choice":

		return fmt.Sprintf(`item.fields.%s ?? ""`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), ""
	case "MultiChoice":
		return fmt.Sprintf(`item.fields.%s ?? []`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), ""
	case "Calculated":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), ""
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
		log.Println("Unknown type", field.Type)
	}

	return "", "", ""
}

func formCode(listName string, formFields string) string {
	return fmt.Sprintf(`

	
	"use client"
	import * as React from "react"
	
	import { useEffect, useState } from "react"
	import { zodResolver } from "@hookform/resolvers/zod"
	
	import { useForm } from "react-hook-form"
	import * as z from "zod"
	
	import {schema} from "."
	import { IProgressProps, ProcessStatusOverlay } from "../../../../../components/progress"
	
	
	import {
		Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage
	} from "../../../../../registry/new-york/ui/form"
	
	import { Input } from "../../../../../registry/new-york/ui/input"
	
	
	// const profileFormSchema = z.object({
	// 	country: z.string({ required_error: "Please select a country." }),
	// 	unit: z.string({
	// 		required_error: "Please select an unit to display.",
	// 	}),
	// 	channels: z.array(z.string({})).optional(),
	// })
	
	type ProfileFormValues = z.infer<typeof schema>
	
	
	
	export function ProfileForm<T>(props: {
		item: T
	
	}) {
	
		const [processing, setProcessing] = useState(false)
		const [processPercentage, setProcessPercentage] = useState(0)
		const [processTitle, setProcessTitle] = useState("")
		const [processDescription, setProcessDescription] = useState("")
		const [lastResult, setlastResult] = useState<any>()
	
	
	
	
	
	
		// This can come from your database or API.
		const defaultValues: Partial<ProfileFormValues> = {
			Title:"xx"
		}
		const form = useForm<ProfileFormValues>({
			resolver: zodResolver(schema),
			defaultValues,
			mode: "onChange",
		})
	
		async function onSubmit(data: ProfileFormValues) {
			setProcessTitle("Saving profile")
			setProcessDescription("Please wait while we save your profile.")
			setProcessPercentage(0)
			setProcessing(true)
	
	
		}
	
	
	
		useEffect(() => {
	
		}, [])
	
	
	
		return (
			<div className="flex">
				<Form {...form}>
					<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
						%s
					</form>
				</Form>
				<ProcessStatusOverlay
					done={!processing}
					title={processTitle}
					description={processDescription}
					progress={processPercentage}
				/>
	
			</div>
		)
	
	
	}
	
	
	
	
			
		



		
	`, formFields)
}
func pageCode(listName string, formFields string) string {
	return fmt.Sprintf(`

	
	"use client"

	import * as React from "react"
	
	import {schema} from "."
	import {ProfileForm} from "./form"
	
	
	
	
	
	export default function ItemPage() {
	
	
	
	
		return (
			<ProfileForm item={undefined}/>
		)
	
	
	}
	
	
	
	
			
		



		
	`)
}
func OutTS(projDir string, list ListInstance) string {

	itemsMap := fmt.Sprintf(`
	export function mapLookup(listName:string,item:any) {
	}
	export function mapLookupMulti(listName:string,item:any) {
	}
// %s
	export function map(item:any) {
	
	
	return {
		Id : item.id,
	Title : item.fields.Title,
	CreatedBy : item.createdBy.user.email,
	Created :new Date(item.createdDateTime),
	ModifiedBy : item.lastModifiedBy.user.email,
	Modified : new Date(item.lastModifiedDateTime),	
		`, list.Title)

	zodMaps := `
	export const schema = z.object({
		CreatedBy : z.string(),
		Created: z.date(),
		ModifiedBy : z.string(),
		Modified: z.date(),
		Id: z.string(),
		Title: z.string(),
		
		`

	formFields := ``
	var standardFields []string = []string{"Id", "Title", "CreatedBy", "Created", "ModifiedBy", "Modified"}
	var fieldNames []string = []string{}
	var dependencies []string = []string{}
	for _, field := range standardFields {
		fieldNames = append(fieldNames, fmt.Sprintf("\"%s\"", field))
	}
	for _, field := range list.Fields.Field {
		fieldMap, zodMap, formField := GetTsField(field)
		if field.Type == "Lookup" || field.Type == "LookupMulti" {
			dependencies = append(dependencies, fmt.Sprintf(`"%s"`, strings.ReplaceAll(strings.ReplaceAll(field.List, "{listid:", ""), "}", "")))
		}
		fieldNames = append(fieldNames, fmt.Sprintf("\"%s\"", field.Name))
		if fieldMap != "" {
			itemsMap += fmt.Sprintf(`%s: %s,
			`, field.Name, fieldMap)

		}
		if zodMap != "" {
			zodMaps += fmt.Sprintf(`%s,
			`, zodMap)

		}

		formFields += formField

	}
	itemsMap += "}}"
	zodMaps += "})"
	sharepointMap := ""

	sharepointMap += fmt.Sprintf(`
	import z from "zod"

		export const listName = "%s"
		`, list.Title)

	sharepointMap += fmt.Sprintf(`export type FieldNames = %s
	`, strings.Join(fieldNames, "|"))
	sharepointMap += fmt.Sprintf(`export const dependencies =[%s]
	`, strings.Join(dependencies, ","))
	sharepointMap += itemsMap

	sharepointMap += zodMaps
	//sharepointMap += "}"
	listName := strings.ReplaceAll(list.Title, " ", "")
	pagedir := createSubdir(projDir, listName)

	writeFile(pagedir, fmt.Sprintf("index.ts"), sharepointMap)
	writeFile(pagedir, fmt.Sprintf("form.tsx"), formCode(listName, formFields))

	writeFile(pagedir, fmt.Sprintf("page.tsx"), pageCode(listName, ""))
	return sharepointMap

}

func Pnp2Ts(sitename string, filename string) string {

	projDir := createProject(sitename) //time.Now().Format("2006-01-02T15:04:05"))
	dollarDir := createSubdir(projDir, "$")
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
	sharepointMap := `
	// ************************************************************
	// Generated by pnp2ts - do not edit
	// www.koksmat.com
	// ************************************************************
	import { https, httpsGetAll } from "@/lib/httphelper"
	import z from "zod"

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
		sharepointMap += OutTS(projDir, list)
	}
	mapSrc := `
	// ************************************************************
	// Generated by pnp2ts - do not edit
	// www.koksmat.com
	// v0.1.0
	// ************************************************************
	

	
	
	
	`
	linkSrc := `
	// ************************************************************
	// Generated by pnp2ts - do not edit
	// www.koksmat.com
	// v0.1.0
	// ************************************************************

	import Link from "next/link";

	


	export default function Page() {
	
		return (<div>
	`

	for _, list := range template.Templates.ProvisioningTemplate.Lists.ListInstance {
		listName := strings.ReplaceAll(list.Title, " ", "")

		linkSrc += fmt.Sprintf(`
<div>
		<Link href="./%s">%s</Link>		
</div>
`, listName, list.Title)

		mapSrc += fmt.Sprintf(`
import { listName as %sList,dependencies as %sDependencies } from "./%s";`, listName, listName, list.Title)
	}

	mapSrc += `
export function map(){

	const dependencies = []
`

	for _, list := range template.Templates.ProvisioningTemplate.Lists.ListInstance {
		listName := strings.ReplaceAll(list.Title, " ", "")
		mapSrc += fmt.Sprintf(`
dependencies.push({listName : %sList,
dependencies:  %sDependencies }) `, listName, listName)
	}
	mapSrc += `

	return dependencies
}

`
	linkSrc += `

	</div>)
}

`
	writeFile(projDir, "index.ts", mapSrc)
	writeFile(dollarDir, "page.tsx", linkSrc)
	return sharepointMap

}
