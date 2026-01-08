package main_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	go_specs_greet "github.com/adettinger/learn-go-with-tests-ATs"
	"github.com/adettinger/learn-go-with-tests-ATs/specifications"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestSimpleContainer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "nginx:latest",
		ExposedPorts: []string{"80/tcp"},
		WaitingFor:   wait.ForHTTP("/").WithPort("80/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start container: %s", err)
	}
	defer container.Terminate(ctx)

	t.Log("Nginx container started successfully")
}

func TestGreeterServer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	os.Setenv("TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE", "npipe:////./pipe/docker_engine")
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")

	t.Log("Starting container creation...")
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../.",
			Dockerfile: "./cmd/httpserver/Dockerfile",
			// set to false if you want less spam, but this is helpful if you're having troubles
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
		// Add these options to keep the container running and capture logs
		AlwaysPullImage: false,
		SkipReaper:      true, // Prevents the container from being removed immediately on exit
		AutoRemove:      false,
	}

	t.Log("Creating container...")
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		// Try to get logs even if container failed to start properly
		if container != nil {
			logReader, logErr := container.Logs(ctx)
			if logErr == nil {
				defer logReader.Close()
				logs, _ := io.ReadAll(logReader)
				t.Logf("Container logs before failure: %s", string(logs))
			} else {
				t.Logf("Failed to get container logs: %s", logErr)
			}
		}
		t.Fatalf("Failed to create/start container: %s", err)
	}
	t.Log("Container created successfully")

	// Get container ID for debugging
	containerID := container.GetContainerID()
	t.Logf("Container ID: %s", containerID)

	cleanupCtx := context.Background()
	t.Cleanup(func() {
		t.Log("Terminating container...")
		terminateCtx, terminateCancel := context.WithTimeout(cleanupCtx, 10*time.Second)
		defer terminateCancel()

		if err := container.Terminate(terminateCtx); err != nil {
			t.Logf("Warning: Failed to terminate container: %s", err)
			// Continue anyway, as this is just cleanup
		}
	})

	// Get the mapped host port instead of assuming it's on localhost:8080
	t.Log("Getting container host...")
	hostIP, err := container.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get container host: %s", err)
	}

	t.Log("Getting mapped port...")
	mappedPort, err := container.MappedPort(ctx, "8080")
	if err != nil {
		t.Fatalf("Failed to get mapped port: %s", err)
	}

	baseURL := fmt.Sprintf("http://%s:%s", hostIP, mappedPort.Port())
	t.Logf("Server URL: %s", baseURL)

	t.Log("Running specification tests...")
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	driver := go_specs_greet.Driver{BaseURL: baseURL, Client: &client}
	// driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	t.Run("server greet specification", func(t *testing.T) {
		specifications.GreetSpecification(t, driver)
	})
}
