
while($true) {
    kubectl port-forward services/prod-mongos 27017:27017 -n percona
    sleep 5
}
  