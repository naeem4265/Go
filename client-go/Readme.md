# Running the API Server from In Cluster

Follow these steps to run the `api-server` within a Kubernetes cluster:

## Step 1: Start a Cluster with a Config File

```shell
$ kind create cluster --config=./kind-config.yaml

## Step 2: Deploy Ingress Controller (nginx)
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

## Step 3: Create Role and Role Binding
Create role and role-binding YAML files, then apply them.
$ kubectl apply -f cluster-role.yaml 
$ kubectl apply -f cluster-role-binding.yaml 
