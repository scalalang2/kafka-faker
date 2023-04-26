package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTemplate(t *testing.T) {
	tpl := Template(`hello {{ personName }}, this is int(10, 100): {{ intBetween 10 100 }}`)
	ret, err := ParseTemplate(tpl)
	require.NoError(t, err)
	t.Logf("generated tpl: %s", ret)
}
