package scaffold

import (
	"fmt"
	"log"
	"strings"
)

func TableTextColumn(name string, displayName string) string {
	return fmt.Sprintf(`	{
		id: "%s",
		accessorKey: "%s",
		header: ({ column }) => (
		  <DataTableColumnHeader column={column} title="%s" />
		),
		cell: ({ row }) => {
		return <div >{row.original.%s}</div>},
		enableSorting: true,
		enableHiding: true,
	  },`, name, name, displayName, name)
}

func TableNumberColumn(name string, displayName string) string {
	return fmt.Sprintf(`	{
		id: "%s",
		accessorKey: "%s",
		header: ({ column }) => (
		  <DataTableColumnHeader column={column} title="%s" />
		),
		cell: ({ row }) => {
		return <div >{row.original.%s}</div>},
		enableSorting: true,
		enableHiding: true,
	  },`, name, name, displayName, name)
}

func TableDateColumn(name string, displayName string) string {
	return fmt.Sprintf(`	{
		id: "%s",
		accessorKey: "%s",
		header: ({ column }) => (
		  <DataTableColumnHeader column={column} title="%s" />
		),
		cell: ({ row }) => {
		return <div >{row.original.%s.toLocaleDateString()}</div>},
		enableSorting: true,
		enableHiding: true,
	  },`, name, name, displayName, name)
}
func FormTextfield(name string, displayName string, description string) string {
	return fmt.Sprintf(`<FormField 
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
		/>`, name, displayName, description)
}

func FormDatefield(name string, displayName string, description string) string {
	return fmt.Sprintf(`<FormField 
			control={form.control}
			name="%s"
			render={({ field }) => (
				<FormItem>
					<FormLabel>%s</FormLabel>
					<FormControl>
						<Input placeholder="" {...field} value={field.value?.toISOString()} />
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}
		/>`, name, displayName, description)
}
func GetTsField(field Field) (string, string, string, string) {
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
						<Input placeholder="" {...field} value={field.value??"" }/>
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}
		/>`, field.Name, field.DisplayName, field.Description), TableTextColumn(field.Name, field.DisplayName)
	case "Note":
		return fmt.Sprintf(`item.fields.%s ? item.fields.%s : ""`, field.Name, field.Name),
			fmt.Sprintf(`%s : z.string()`, field.Name), fmt.Sprintf(`<FormField
			control={form.control}
			name="%s"
			render={({ field }) => (
				<FormItem>
					<FormLabel>%s</FormLabel>
					<FormControl>
						<Input placeholder="" {...field} value={field.value??""} />
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}
		/>`, field.Name, field.DisplayName, field.Description), ""
	case "Number":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.number()`, field.Name), fmt.Sprintf(`<FormField
			control={form.control}
			name="%s"
			render={({ field }) => (
				<FormItem>
					<FormLabel>%s</FormLabel>
					<FormControl>
						<Input placeholder="" {...field} value={field.value??""}/>
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}
		/>`, field.Name, field.DisplayName, field.Description), TableNumberColumn(field.Name, field.DisplayName)
	case "Boolean":
		return fmt.Sprintf(`item.fields.%s ? true : false`, field.Name),
			fmt.Sprintf(`%s : z.boolean()`, field.Name), "", ""
	case "DateTime":
		return fmt.Sprintf(`new Date(item.fields.%s)`, field.Name),
			fmt.Sprintf(`%s : z.date()`, field.Name), fmt.Sprintf(`<FormField
			control={form.control}
			name="%s"
			render={({ field }) => (
				<FormItem>
					<FormLabel>%s</FormLabel>
					<FormControl>
						<Input placeholder="" {...field} value={field.value?.toISOString()??""} />
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}
		/>`, field.Name, field.DisplayName, field.Description), TableDateColumn(field.Name, field.DisplayName)
	case "User":
		return fmt.Sprintf(`item.fields.%s ?? ""`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), fmt.Sprintf(`<FormField
			control={form.control}
			name="%s"
			render={({ field }) => {
			
				return (
				<FormItem>
					<FormLabel>%s</FormLabel>
					<FormControl>
						<Input placeholder="" {...field}  value={field.value??""}/>
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}}
		/>`, field.Name, field.DisplayName, field.Description), ""
	case "UserMulti":
		return fmt.Sprintf(`item.fields.%s ?? ""`, field.Name),
			fmt.Sprintf(`%s : z.string().array().nullable()`, field.Name), "", ""
	case "Lookup":
		return fmt.Sprintf(`mapLookup(%s,item.fields.%sLookupId)`, strings.ReplaceAll(strings.ReplaceAll(field.List, "{listid:", "\""), "}", "\""),
				field.Name),
			fmt.Sprintf(`%s : z.object({
				LookupId:z.number(),
				LookupValue:z.string()
			  }).nullable()`, field.Name),

			fmt.Sprintf(`<FormField
			  control={form.control}
			  name="%s"
			  render={({ field }) => (
				  <FormItem>
					  <FormLabel>%s</FormLabel>
					  <FormControl>
					  
						  <Input placeholder="" {...field} value={field.value? field.value?.LookupValue:""}/>
					  </FormControl>
					  <FormDescription>
						  %s
					  </FormDescription>
					  <FormMessage />
				  </FormItem>
			  )}
		  />`, field.Name, field.DisplayName, field.Description), ""
	case "LookupMulti":
		return fmt.Sprintf(`mapLookupMulti(%s,item.fields.%sLookupId)`, strings.ReplaceAll(strings.ReplaceAll(field.List, "{listid:", "\""), "}", "\""), field.Name),
			fmt.Sprintf(`%s : z.object({
				LookupId:z.number(),
				LookupValue:z.string()
			  }).array().nullable()`, field.Name), "", ""
	case "URL":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), fmt.Sprintf(`<FormField
			control={form.control}
			name="%s"
			render={({ field }) => (
				<FormItem>
					<FormLabel>%s</FormLabel>
					<FormControl>
						<Input placeholder="" {...field} value={field.value??""}/>
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}
		/>`, field.Name, field.DisplayName, field.Description), ""
	case "Choice":

		return fmt.Sprintf(`item.fields.%s ?? ""`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), fmt.Sprintf(`<FormField
			control={form.control}
			name="%s"
			render={({ field }) => (
				<FormItem>
					<FormLabel>%s</FormLabel>
					<FormControl>
						<Input placeholder="" {...field} value={field.value??""}/>
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}
		/>`, field.Name, field.DisplayName, field.Description), TableTextColumn(field.Name, field.DisplayName)
	case "MultiChoice":
		return fmt.Sprintf(`item.fields.%s ?? []`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), "", ""
	case "Calculated":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.string().nullable()`, field.Name), fmt.Sprintf(`<FormField
			control={form.control}
			name="%s"
			render={({ field }) => (
				<FormItem>
					<FormLabel>%s</FormLabel>
					<FormControl>
						<Input placeholder="" {...field} value={field.value??""}/>
					</FormControl>
					<FormDescription>
						%s
					</FormDescription>
					<FormMessage />
				</FormItem>
			)}
		/>`, field.Name, field.DisplayName, field.Description), ""
		// case "Attachments":
		// 	return "string"
		// case "Guid":
		// 	return "string"
	case "Integer":
		return fmt.Sprintf(`item.fields.%s`, field.Name),
			fmt.Sprintf(`%s : z.number()`, field.Name), fmt.Sprintf(`<FormField
		control={form.control}
		name="%s"
		render={({ field }) => (
			<FormItem>
				<FormLabel>%s</FormLabel>
				<FormControl>
					<Input placeholder="" {...field} value={field.value??""}/>
				</FormControl>
				<FormDescription>
					%s
				</FormDescription>
				<FormMessage />
			</FormItem>
		)}
	/>`, field.Name, field.DisplayName, field.Description), TableNumberColumn(field.Name, field.DisplayName)
	// case "Counter":

	// 	return "number"
	// case "Currency":

	// 	return "number"
	default:
		log.Println("Unknown type", field.Type)
	}

	return "", "", "", ""
}
