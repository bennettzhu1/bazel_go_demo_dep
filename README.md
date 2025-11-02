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

### Cleanup
```bash
helm uninstall my-app
kind delete cluster
```

## Bazel Workflow

### Setup
```bash
# Create cluster
kind create cluster

# Build and load image (AMD64 for Kind)
bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //:load_image
kind load docker-image hello-k8s:latest
```

### Deploy
```bash
# Install chart
bazel run //helm-chart:chart.install
```

### Access
```bash
kubectl port-forward svc/hello-k8s 8080:80
curl http://localhost:8080
```

### Cleanup
```bash
helm uninstall chart
kind delete cluster
```

## Optional: Inspect Bazel Outputs

```bash
# View rendered templates
bazel build //helm-chart:rendered_chart
cat bazel-bin/helm-chart/rendered_chart/hello-k8s/templates/deployment.yaml

# View packaged chart
bazel build //helm-chart:chart
ls bazel-bin/helm-chart/
```

## Useful Commands

```bash
# Check cluster status
kubectl get nodes
kubectl get pods
kubectl get services

# View contexts
kubectl config get-contexts
kubectl config use-context kind-kind

# Debug
kubectl logs <pod-name>
kubectl describe pod <pod-name>
```

## Key Differences

**Traditional**: `docker build` + `helm template | kubectl apply` or `helm install`
**Bazel**: `bazel run :load_image` + `bazel run :chart.install`

Bazel manages the build + packaging. Kubernetes operations (deploy, port-forward) are the same.
