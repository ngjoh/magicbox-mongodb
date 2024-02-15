"use client"

import { useContext, useEffect, useState } from "react"
import { Context } from "./context"
import { Button } from "@/components/ui/button"
import { set } from "date-fns"
import { useProcess } from "@/koksmat/useprocess"
import { APPNAME } from "../appid"

export default function Index() {
  const context = useContext(Context)

  const { isLoading, error, data } = useProcess(
    APPNAME,
    [],
    20,
    "echo"
  )

  return (
    <div>
      <pre>
        {data}

      </pre>

    </div>
  )
}
