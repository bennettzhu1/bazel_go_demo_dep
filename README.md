# Kubernetes Basics Tutorial 🚀

This tutorial demonstrates the fundamentals of Kubernetes and Helm, which will help you understand what `rules_gitops` and `rules_helm` are doing under the hood.

## 📚 What You'll Learn

1. **Docker** - Containerizing a Go application
2. **Kind** - Creating a local Kubernetes cluster
3. **kubectl apply** - Deploying raw Kubernetes manifests
4. **Helm** - Templating and packaging Kubernetes manifests
5. **Port Forwarding** - Accessing services from your local machine

---

## 🛠️ Prerequisites

Install these tools:
```bash
# Install Kind (Kubernetes in Docker)
brew install kind

# Install kubectl (Kubernetes CLI)
brew install kubectl

# Install Helm (Kubernetes package manager)
brew install helm

# Verify installations
kind --version
kubectl version --client
helm version
```

---

## 📖 Tutorial Steps

### Step 1: Create a Kind Cluster

```bash
cd ~/workspace/k8s-basics

# Create a Kind cluster
kind create cluster --name k8s-basics

# Verify it's running
kubectl cluster-info --context kind-k8s-basics
kubectl get nodes
```

**What just happened?**
- Kind created a Docker container that runs Kubernetes
- `kubectl` is now configured to talk to this cluster (check `~/.kube/config`)
- You have a fully functional Kubernetes cluster on your laptop!

---

### Step 2: Build and Load the Docker Image

```bash
# Build the Docker image
docker build -t hello-k8s:latest .

# Load the image into the Kind cluster
# (Kind nodes can't access your local Docker daemon by default)
kind load docker-image hello-k8s:latest --name k8s-basics

# Verify the image is loaded
docker exec -it k8s-basics-control-plane crictl images | grep hello-k8s
```

**What just happened?**
- Built a Docker image with your Go app
- Loaded it into the Kind cluster's internal registry
- The Kubernetes nodes can now use this image

---

### Step 3: Deploy with Raw Kubernetes YAML (kubectl apply)

```bash
# Apply the Deployment
kubectl apply -f kubernetes/deployment.yaml

# Apply the Service
kubectl apply -f kubernetes/service.yaml

# Watch the pods come up
kubectl get pods -w
# (Press Ctrl+C when both pods are Running)

# Check the deployment
kubectl get deployments
kubectl get pods
kubectl get services
```

**What just happened?**
- `kubectl apply` sent the YAML to the Kubernetes API server
- Kubernetes created:
  - A **Deployment** (manages the desired state of 2 pods)
  - 2 **Pods** (running your Go app)
  - A **Service** (stable IP/DNS for accessing the pods)

**View the raw YAML:**
```bash
# See what Kubernetes created
kubectl get deployment hello-k8s -o yaml
kubectl get service hello-k8s -o yaml
```

---

### Step 4: Access the Application (Port Forwarding)

```bash
# Forward local port 8080 to the service
kubectl port-forward service/hello-k8s 8080:80

# In another terminal, test it:
curl http://localhost:8080
# Output: Hello from hello-k8s-raw! 🚀

# Check the logs
kubectl logs -l app=hello-k8s --all-containers=true --tail=10
```

**What just happened?**
- `kubectl port-forward` created a tunnel from your laptop to the Service
- Traffic to `localhost:8080` → forwarded to the Service (port 80) → routed to one of the pods (port 8080)
- This is EXACTLY what `rules_gitops`'s `k8s_test_setup` does!

---

### Step 5: Clean Up Raw Deployment

```bash
# Delete the resources
kubectl delete -f kubernetes/deployment.yaml
kubectl delete -f kubernetes/service.yaml

# Verify they're gone
kubectl get pods
kubectl get services
```

---

### Step 6: Deploy with Helm

#### 6a. Understanding Helm Templates

```bash
# Preview what Helm will generate (without deploying)
helm template hello-k8s ./helm-chart

# Save the rendered YAML to a file
helm template hello-k8s ./helm-chart > /tmp/helm-rendered.yaml

# Compare it to your raw YAML
diff kubernetes/deployment.yaml /tmp/helm-rendered.yaml
```

**What just happened?**
- `helm template` replaced all the `{{ .Values.* }}` placeholders with values from `values.yaml`
- This is EXACTLY what `rules_helm`'s `helm_template` does in Bazel!
- The output is plain Kubernetes YAML (same as what you wrote by hand)

#### 6b. Deploy with Helm Install

```bash
# Install the Helm chart
helm install hello-k8s ./helm-chart

# Check what was created
helm list
kubectl get pods -l app=hello-k8s
kubectl get services
```

**What just happened?**
- `helm install` = `helm template` + `kubectl apply`
- Helm also tracks the release in its database (so you can `helm upgrade`, `helm rollback`, etc.)

#### 6c. Test with Port Forwarding

```bash
# Forward to the Helm-deployed service
kubectl port-forward service/hello-k8s 8080:80

# In another terminal:
curl http://localhost:8080
# Output: Hello from hello-k8s-helm! 🚀
```

#### 6d. Customize with Values

```bash
# Create a custom values file
cat > custom-values.yaml << EOF
replicaCount: 3
appName: hello-k8s-custom
EOF

# Upgrade the release with custom values
helm upgrade hello-k8s ./helm-chart -f custom-values.yaml

# Verify the changes
kubectl get pods -l app=hello-k8s
# (Should now see 3 pods)

curl http://localhost:8080
# Output: Hello from hello-k8s-custom! 🚀
```

**What just happened?**
- Helm merged `custom-values.yaml` with `values.yaml`
- This is EXACTLY what you did with `test-values.yaml` in the traffic-cell migration!

---

### Step 7: Clean Up

```bash
# Uninstall the Helm release
helm uninstall hello-k8s

# Delete the Kind cluster
kind delete cluster --name k8s-basics
```

---

## 🔗 Connecting to rules_gitops & rules_helm

Now you understand the raw tools. Here's how Bazel wraps them:

### `rules_helm` does this:
```python
helm_template(
    name = "hello_k8s_manifests",
    chart = "//helm-chart",
    values = "custom-values.yaml",
)
```
**Equivalent to:**
```bash
helm template hello-k8s ./helm-chart -f custom-values.yaml > manifests.yaml
```

### `rules_gitops` does this:
```python
k8s_deploy(
    name = "deploy",
    manifests = [":hello_k8s_manifests"],
)

k8s_test_setup(
    name = "test_setup",
    kubeconfig = "~/.kube/config",
    objects = [":deploy"],
    portforward_services = [
        "hello-k8s:8080:80",  # local:service
    ],
)
```
**Equivalent to:**
```bash
kubectl apply -f manifests.yaml
kubectl wait --for=condition=Ready pod -l app=hello-k8s
kubectl port-forward service/hello-k8s 8080:80 &
```

---

## 🎯 Key Takeaways

1. **Docker** = Package your app + dependencies into an image
2. **Kind** = Run a real Kubernetes cluster locally (as Docker containers)
3. **Deployment** = Describes desired state (e.g., 2 replicas of hello-k8s)
4. **Service** = Stable network endpoint to access your pods
5. **kubectl apply** = Send YAML to Kubernetes API
6. **Helm** = Templating system to generate YAML dynamically
7. **Port Forwarding** = Tunnel from localhost to a Kubernetes Service

**rules_gitops + rules_helm** = Bazel rules that automate all of this for testing!

---

## 🧪 Experiment Ideas

Try these to deepen your understanding:

1. **Change replica count** in `deployment.yaml` and re-apply
2. **Modify the Service port** and see how port-forward changes
3. **Add a new endpoint** to `main.go` (e.g., `/version`) and redeploy
4. **Create a new Helm template** (e.g., `configmap.yaml`)
5. **Break something** (e.g., remove the Service) and see what happens to port-forward

---

## 📚 Next Steps

Now that you understand the basics, explore:
- [rules_helm documentation](https://github.com/fasterci/rules_helm)
- [rules_gitops documentation](https://github.com/fasterci/rules_gitops)
- Kubernetes concepts: Namespaces, ConfigMaps, Secrets, RBAC
- Helm advanced features: Hooks, Tests, Dependencies

Happy learning! 🎉


