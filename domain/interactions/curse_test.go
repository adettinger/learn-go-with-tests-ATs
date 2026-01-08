package interactions_test

import (
	"testing"

	"github.com/adettinger/learn-go-with-tests-ATs/domain/interactions"
	"github.com/adettinger/learn-go-with-tests-ATs/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(interactions.Curse),
	)
}
