## Overview

K8sWatcher is a simple kubernetes operator with which you can monitor and get alerts about:
 - pod healty
 
 
## Quick Start

### Install

```bash
kubectl apply -f deploy/crds/sbtech_v1_k8swatchernotifier_crd.yaml
kubectl apply -f deploy/cluster_role_view.yaml
kubectl apply -f deploy/role.yaml
kubectl apply -f deploy/role_binding.yaml
kubectl apply -f deploy/service_account.yaml
kubectl apply -f deploy/operator.yaml
kubectl create clusterrolebinding k8s-watcher-view --clusterrole=view --serviceaccount=monitoring:k8s-watcher
```
Modify and apply k8swatcher custom resource in any namespaces in which you need to monitor pods

```bash
kubectl apply -f deploy/crds/sbtech_v1_k8swatchernotifier_cr.yaml -n namespaces
```

Enjoy :)








