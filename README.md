
### Installing kubernetes

* Run the `k8s` role on all nodes

* Run `playbooks/utils/k8s/install.yaml`, this will spin up the cluster initially

* Manifest commands that need to be applied manually for CNI and thus the cluster as a whole to work:

```bash
# Install calico
kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.25.0/manifests/calico.yaml
```
