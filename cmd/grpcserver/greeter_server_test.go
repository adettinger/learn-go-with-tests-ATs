package main_test

import (
	"fmt"
	"testing"

	"github.com/adettinger/learn-go-with-tests-ATs/adapters"
	"github.com/adettinger/learn-go-with-tests-ATs/adapters/grpcserver"
	"github.com/adettinger/learn-go-with-tests-ATs/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerFilePath = "./cmd/grpcserver/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, &driver)
}
