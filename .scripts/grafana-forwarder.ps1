


while($true) {
    kubectl port-forward services/kube-prometheus-grafana 3000:80 -n monitoring 
    sleep 5
}
  