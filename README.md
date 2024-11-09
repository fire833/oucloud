
### Installing kubernetes

* Run the `k8s` role on all nodes

* Run `playbooks/utils/k8s/install.yaml`, this will spin up the cluster initially

* Manifest commands that need to be applied manually for CNI and thus the cluster as a whole to work:

```bash
# Install the tigera operator.
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.26.1/manifests/tigera-operator.yaml
# Install the required Calico CRD. You may want to switch this out for a custom 
# manifest which will include a custom CIDR range.
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.26.1/manifests/custom-resources.yaml
kubectl create -f https://calico-v3-25.netlify.app/archive/v3.26.1/manifests/calico.yaml
```

