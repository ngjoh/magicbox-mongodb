package scaffold

import "fmt"

func itemPageCode(listName string, formFields string) string {
	return fmt.Sprintf(`
	"use client"

	import * as React from "react"
	import { Button } from "@/registry/new-york/ui/button"
	import { MagicboxContext } from "@/app/magicbox-context"
	import {
	  useSharePointList,
	  useSharePointListItem,
	  version,
	  
	} from "@/app/sharepoint"
	
	import { ItemType, dependencies, listName, schema,map,listURL } from ".."
	import { ItemForm } from "../form"
	import Link from "next/link"

	export default function ItemPage({
	  params,
	}: {
	  params: { itemid: string; site: string }
	}) {
	  const magicbox = React.useContext(MagicboxContext)
	  const { site, itemid } = params
	
	  const { item,error,isLoading ,itemRaw} = useSharePointListItem<ItemType>(
		magicbox.session?.accessToken ?? "",
		magicbox.tenant,
		site,
		listName,
		itemid,
		map,
		schema,
	  )
	
	  return (
		<div>
		<div className="container"  style={{minHeight:"100vh"}}>
		{error && <div className="text-red-600">{error}</div>}	
		{isLoading && <div className="text-green-600">Loading...</div>}
		<div className="rounded-xl bg-slate-800 text-gray-50">
		<div className="p-4 text-3xl">{item?.Title} </div>

</div>
<div className="flex">
<div className="pl-5 text-xs">List: {listName} | Created: {item?.Created?.toLocaleDateString()} by {item?.CreatedBy} | Modified: {item?.Modified?.toLocaleDateString()} by {item?.ModifiedBy} </div>
<div className="grow" />
<Button variant={"link"}><Link target="_blank" href={"https://" + magicbox.tenant + ".sharepoint.com/sites/" + site + "/"+ listURL + "/DispForm.aspx?ID="+itemid}>View in SharePoint</Link></Button>
</div>
<ItemForm item={item} />
		  <pre>
				{JSON.stringify(item, null, 2)}
			</pre>
			<pre >
				{JSON.stringify(itemRaw, null, 2)}
			</pre>
			</div>
		</div>
	  )

	}
	
	`)
}
