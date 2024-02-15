"use client"

import { useContext, useEffect, useState } from "react"

import { useProcess } from "@/koksmat/useprocess"


import { Context, ContextProps, Options } from "./context"
import { IView } from "./components/viewcards"
import Script from "next/script"


type Props = {
  children?: React.ReactNode
  rootPath: string

  isLocalEnv: boolean
}

export const ContextProvider = ({ children, rootPath, isLocalEnv }: Props) => {


  const [options, setoptions] = useState<Options>({
    showContext: true,
    showNavigation: false,
  })

  const [isloaded, setisloaded] = useState(false)
  const [workspaces, setworkspaces] = useState<IView[]>([])
  const [loadedOffice, setloadedOffice] = useState(false)
  useEffect(() => {
    const load = async () => {

      setisloaded(true)
    }
    load()

  }, [])

  const koksmat: ContextProps = {

    isloaded,
    workspaces,
    options,
    setOptions: function (changes: Options): void {
      setoptions({ ...options, ...changes })
    },
    rootPath
  }

  return <Context.Provider value={koksmat}>
    {children}</Context.Provider>
}
