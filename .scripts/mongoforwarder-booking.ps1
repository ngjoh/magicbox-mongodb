
while($true) {
    kubectl port-forward services/booking-mongos 27018:27017 -n percona
    sleep 5
}
  