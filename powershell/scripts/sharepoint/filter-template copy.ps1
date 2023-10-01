
$file = "allpages-template"
$filename = "$PSScriptRoot/$file.xml"


$filenameOut = "$PSScriptRoot/$file-filtered.xml"
[xml]$xml = Get-Content $filename 



$nsmgr = New-Object System.Xml.XmlNamespaceManager $xml.NameTable
$nsmgr.AddNamespace('pnp','http://schemas.dev.office.com/PnP/2022/09/ProvisioningSchema')


$XPath = "//pnp:ProvisioningTemplate"

$ns = New-Object System.Xml.XmlNamespaceManager($xml.NameTable)
$product = $xml.SelectSingleNode($XPath, $nsmgr)
function RemoveNode($path) {
    $node = $xml.SelectSingleNode( $path, $nsmgr)
    if ($node  -ne $null) {
        $product.RemoveChild($node)
    }
}

RemoveNode("//pnp:PropertyBagEntries")
RemoveNode("//pnp:Security")
RemoveNode("//pnp:SiteFields")
RemoveNode("//pnp:Lists")
#RemoveNode("//pnp:Features")
RemoveNode("//pnp:Footer")

RemoveNode("//pnp:Navigation")
RemoveNode("//pnp:WebSettings")
RemoveNode("//pnp:SiteSettings")
RemoveNode("//pnp:RegionalSettings")
RemoveNode("//pnp:CustomActions")
#RemoveNode("//pnp:ApplicationLifecycleManagement")
$pages = $xml.SelectNodes("//pnp:ClientSidePages/pnp:ClientSidePage", $nsmgr)
foreach ($page in $pages) {
    $pageFileName =  $page.GetAttribute("PageName")
    if (!($pageFileName.ToLower().StartsWith("germany"))) {
        $page.ParentNode.RemoveChild($page)
        write-host "Removing page ", $pageFileName
    }
    else {
        write-host "Keeping page ", $pageFileName
    }
    $page.SetAttribute("Overwrite", "true")
}

write-host "Number of pages ",$pages.Count
#$nav = $xml.SelectSingleNode( "//pnp:CurrentNavigation/pnp:StructuralNavigation", $nsmgr)
#$nav.SetAttribute("RemoveExistingNodes","true")
$xml.Save($filenameOut)

