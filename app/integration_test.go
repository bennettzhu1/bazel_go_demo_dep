package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/fasterci/rules_gitops/client"
)

var setup client.K8STestSetup

func TestMain(m *testing.M) {
	setup = client.K8STestSetup{
		WaitForPods: []string{"hello-k8s"},
		PortForwardServices: map[string]int{
			"hello-k8s": 8080,
		},
	}
	setup.TestMain(m)
}

func TestSimpleServer(t *testing.T) {
	appServerPort := setup.GetServiceLocalPort("hello-k8s")
	response, err := http.Get(fmt.Sprintf("http://localhost:%d", appServerPort))
	if err != nil {
		t.Fatalf("Failed to connect to app server: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	body, _ := io.ReadAll(response.Body)
	if !strings.Contains(string(body), "Hello from hello-k8s-helm") {
		t.Errorf("Unexpected content returned: %q", string(body))
	}
}
