# Kubernetes + Bazel Basics

## Prerequisites

```bash
brew install kind kubectl helm
```

## Traditional Workflow

### Setup
Works if helm-chart references hello-k8s:latest, we changed it to 
```bash
kind create cluster

docker build -t hello-k8s:latest .
kind load docker-image hello-k8s:latest
```

### Deploy
```bash
# helm template won't work since we changed 
# helm template ./manifests | kubectl apply -f -
kubectl apply -f ./manifests
```

### Access
```bash
kubectl port-forward svc/hello-k8s 8080:80
curl http://localhost:8080
```


## Bazel Workflow

### Using helm_template and kubectl
```bash
bazel run //app:load_image
kind load docker-image hello-k8s:latest
bazel build //helm-chart:rendered_chart
kubectl apply -f bazel-bin/helm-chart/rendered_chart.yaml
```

### Use k8s_deploy (no need to `load_image` and `kind load`)
```bash
kind delete cluster
./setup.sh
```
or
```bash
./create_kind_cluster.sh
kubectl create namespace default
bazel run //app:deploy.apply
```

./create_kind_cluster sets up a registry with kind-registry:5000, and then the kind cluster gets created with a special config that substitutes pushes to localhost:5000 to the kind cluster to bypass the need to oci_load and kind load.

We run `deploy.apply` (`k8s_deploy.apply` ) which pushes the image to `localhost:5000` (local docker registry), and kind is configured to pull there bypassing manual kind load. It's available to the template, which uses a special `{{//app:image_index}}`" to reference it in the deployment
