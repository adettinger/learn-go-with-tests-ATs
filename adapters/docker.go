package adapters

import (
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartDockerServer(
	t testing.TB,
	port string,
	dockerFilePath string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	t.Helper()

	// Env variables for running testcontainers-go in windows
	os.Setenv("TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE", "npipe:////./pipe/docker_engine")
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")

	t.Log("Starting container creation...")
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    "../../.",
			Dockerfile: dockerFilePath,
			// set to false if you want less spam, but this is helpful if you're having troubles
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
		// options to keep the container running and capture logs
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
		if container != nil {
			captureContainerLogsToTest(t, container, ctx)
		}
		t.Fatalf("Failed to create/start container: %s", err)
	}
	t.Log("Container created successfully")
	t.Logf("Container ID: %s", container.GetContainerID())

	t.Cleanup(func() {
		t.Log("Terminating container...")
		if err := container.Terminate(context.Background()); err != nil {
			t.Fatalf("Warning: Failed to terminate container: %s", err)
		}
	})
}

func captureContainerLogsToTest(t testing.TB, container testcontainers.Container, ctx context.Context) {
	t.Helper()
	logReader, logErr := container.Logs(ctx)
	if logErr == nil {
		defer logReader.Close()
		logs, _ := io.ReadAll(logReader)
		t.Logf("Container logs before failure: %s", string(logs))
	} else {
		t.Logf("Failed to get container logs: %s", logErr)
	}
}
