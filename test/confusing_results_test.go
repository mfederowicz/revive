package test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

func TestConfusingResults(t *testing.T) {
	testRule(t, "confusing_results", &rule.ConfusingResultsRule{})
}
