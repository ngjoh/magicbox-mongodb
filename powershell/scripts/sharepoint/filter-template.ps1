$filename = "$PSScriptRoot/template.xml"
$filenameOut = "$PSScriptRoot/template-filtered.xml"
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

#RemoveNode("//pnp:CustomActions")
#RemoveNode("//pnp:ApplicationLifecycleManagement")

$nav = $xml.SelectSingleNode( "//pnp:CurrentNavigation/pnp:StructuralNavigation", $nsmgr)
$nav.SetAttribute("RemoveExistingNodes","true")
$xml.Save($filenameOut)

