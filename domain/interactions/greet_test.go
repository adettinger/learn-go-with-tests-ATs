package go_specs_greet_test

import (
	"testing"

	go_specs_greet "github.com/adettinger/learn-go-with-tests-ATs/domain/interactions"
	"github.com/adettinger/learn-go-with-tests-ATs/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(go_specs_greet.Greet),
	)
}
