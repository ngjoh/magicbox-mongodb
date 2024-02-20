$ns = "percona"
$spaces = kubens

if (!$spaces -contains $ns) {
    kubectl create namespace $ns
}
$current = kubens -c
if ($current -ne $ns) {
    kubens $ns
}
kubectx
kubens