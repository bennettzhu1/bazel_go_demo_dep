# Kubernetes + Bazel Basics

## Prerequisites

```bash
brew install kind kubectl helm
```

## Traditional Workflow

### Setup
```bash
# Create cluster
kind create cluster

# Build and load image
docker build -t hello-k8s:latest .
kind load docker-image hello-k8s:latest
```

### Deploy
```bash
# Render templates and apply
helm template my-app ./helm-chart | kubectl apply -f -

# Or install directly
helm install my-app ./helm-chart
```

### Access
```bash
kubectl port-forward svc/hello-k8s 8080:80
curl http://localhost:8080
```


## Bazel Workflow

### Using helm_template
```bash
bazel build //helm-chart:rendered_chart
kubectl apply -f bazel-bin/helm-chart/rendered_chart.yaml
```

### Integ
```bash
kind delete cluster
./setup.sh
```
or run it after creating a cluster and a namespace
