


while($true) {
    kubectl port-forward services/kube-prometheus-kube-prome-prometheus 9090:9090 -n monitoring
    sleep 5
}
  