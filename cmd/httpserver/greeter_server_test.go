package main_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/adettinger/learn-go-with-tests-ATs/adapters"
	"github.com/adettinger/learn-go-with-tests-ATs/adapters/httpserver"
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
	var (
		port    = "8080"
		baseURL = fmt.Sprintf("http://localhost:%s", port)
		driver  = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, "httpserver", "")

	t.Run("server greet specification", func(t *testing.T) {
		specifications.GreetSpecification(t, driver)
	})
	t.Run("server curse specification", func(t *testing.T) {
		specifications.CurseSpecification(t, driver)
	})
}
