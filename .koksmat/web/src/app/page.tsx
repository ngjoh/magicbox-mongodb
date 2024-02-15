
import { redirect } from 'next/navigation'
import {APPNAME} from "@/app/appid"
export default function Home() {
  redirect("/"+APPNAME)
  return (
   
   <div>

   </div>  );
}
