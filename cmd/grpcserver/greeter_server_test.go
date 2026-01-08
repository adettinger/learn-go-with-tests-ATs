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
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, "grpcserver", "")
	specifications.GreetSpecification(t, &driver)
}
