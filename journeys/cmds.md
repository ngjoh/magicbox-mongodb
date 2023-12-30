```bash
kubectl exec --namespace magicbox-miller -it etcd-client -- bash
etcdctl --user root:$ROOT_PASSWORD get "" --prefix --keys-only
etcdctl --user root:$ROOT_PASSWORD get "Purpose@cava:fbe8f4ce-f80e-421a-b9ab-3e0aa29db4b4" --print-value-only

```