package main_test

import (
	"fmt"
	"testing"

	"github.com/adettinger/learn-go-with-tests-ATs/adapters"
	"github.com/adettinger/learn-go-with-tests-ATs/adapters/grpcserver"
	"github.com/adettinger/learn-go-with-tests-ATs/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)
	t.Cleanup(driver.Close)
	adapters.StartDockerServer(t, port, "grpcserver", "")
	specifications.GreetSpecification(t, &driver)
	specifications.CurseSpecification(t, &driver)
}
