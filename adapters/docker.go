package adapters

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartDockerServer(
	t testing.TB,
	port string,
	binToBuild string,
	imageName string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	t.Helper()

	// Env variables for running testcontainers-go in windows
	os.Setenv("TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE", "npipe:////./pipe/docker_engine")
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")

	t.Log("Starting container creation...")
	req := testcontainers.ContainerRequest{
		ExposedPorts: []string{fmt.Sprintf("%s:%s", port, port)},
		WaitingFor:   wait.ForListeningPort(nat.Port(port)).WithStartupTimeout(5 * time.Second),
		// options to keep the container running and capture logs
		AlwaysPullImage: false,
		SkipReaper:      true, // Prevents the container from being removed immediately on exit
		AutoRemove:      false,
	}
	if imageName == "" {
		req.FromDockerfile = testcontainers.FromDockerfile{
			Context: "../../.",
			Dockerfile: "Dockerfile",
			// set to false if you want less spam, but this is helpful if you're having troubles
			BuildArgs: map[string]*string{
				"bin_to_build":   &binToBuild,
				"port_to_expose": &port,
			},
			PrintBuildLog: true,
		}
	} else {
		req.Image = imageName
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

	cleanupCtx := context.Background()
	t.Cleanup(func() {
		t.Log("Terminating container...")
		terminateCtx, terminateCancel := context.WithTimeout(cleanupCtx, 10*time.Second)
		defer terminateCancel()

		if err := container.Terminate(terminateCtx); err != nil {
			t.Logf("Warning: Failed to terminate container: %s", err)
			// Continue anyway incase cleanup is still running
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
