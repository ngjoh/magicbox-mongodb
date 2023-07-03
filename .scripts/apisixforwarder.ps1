


while($true) {
    kubectl port-forward services/apisix-dashboard 9000:80 -n gateway
    sleep 5
}
  