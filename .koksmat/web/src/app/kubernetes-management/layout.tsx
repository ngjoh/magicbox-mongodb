

import { ContextProvider } from "./contextprovider"
import { KoksmatProvider } from "@/koksmat/contextprovider"
import {APPNAME} from "@/app/appid"


export default function JourneyLayoutRoot(props: {
    children: React.ReactNode
}) {


    return (
        // <TopNav rootPath="/officeaddin/" />
        <KoksmatProvider app={APPNAME} hasTopnav={false}>
            <ContextProvider rootPath={""} isLocalEnv={false}>
                <div>
                    {props.children}
                </div>
            </ContextProvider>
        </KoksmatProvider>

    )
}