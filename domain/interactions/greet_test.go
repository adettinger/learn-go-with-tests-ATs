package interactions_test

import (
	"testing"

	"github.com/adettinger/learn-go-with-tests-ATs/domain/interactions"
	"github.com/adettinger/learn-go-with-tests-ATs/specifications"
	"github.com/alecthomas/assert/v2"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)

	t.Run("default name to world", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
