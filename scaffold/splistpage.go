package scaffold

import "fmt"

func listPageCode(listName string, formFields string) string {
	return fmt.Sprintf(`

	"use client"

import * as React from "react"
import { useEffect, useState } from "react"
import Link from "next/link"
import { ColumnDef } from "@tanstack/react-table"

import { DataTableColumnHeader } from "@/components/table/components/data-table-column-header"
import { GenericItem } from "@/components/table/data/schema"
import { Button } from "@/registry/new-york/ui/button"
import { MagicboxContext } from "@/app/magicbox-context"
import { useSharePointList } from "@/app/sharepoint"

import { ItemType, listName, map } from "."
import { %sTable } from "./table"

export default function ItemPage({ params }: { params: { site: string } }) {
  const magicbox = React.useContext(MagicboxContext)
  const { site } = params
  const tenant = magicbox.tenant
  const { items, error, isLoading } = useSharePointList(
    magicbox.session?.accessToken ?? "",
    tenant,
    site,
    listName
  )
  const [parsedItems, setparsedItems] = useState<ItemType[]>([])

  useEffect(() => {
    setparsedItems(items.map((item: any) => map(item)))
  }, [items])

  

  return (
    <div className="container" style={{ minHeight: "100vh" }}>
      {isLoading && <div>Loading...</div>}
      {error && <div className="text-red-700">{error}</div>}
      <%sTable
        items={parsedItems}
        site={site}
        listName={listName}
      />
    </div>
  )
}

		
	`, listName, listName)
}
