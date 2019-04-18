package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInArgs(t *testing.T) {
	tt := []struct {
		name       string
		args       []string
		table      string
		assertions func(*testing.T, []string, string)
	}{
		{
			name:  "intersecting values",
			args:  []string{"TableA", "TableB", "TableC"},
			table: "TableB",
			assertions: func(t *testing.T, a []string, tb string) {
				b := inArgs(a, tb)
				assert.True(t, b)
			},
		},
		{
			name:  "non intersecting values",
			args:  []string{"TableA", "TableB", "TableC"},
			table: "TableD",
			assertions: func(t *testing.T, a []string, tb string) {
				b := inArgs(a, tb)
				assert.False(t, b)
			},
		},
		{
			name:  "empty args",
			args:  []string{},
			table: "TableA",
			assertions: func(t *testing.T, a []string, tb string) {
				b := inArgs(a, tb)
				assert.True(t, b)
			},
		},
	}

	for _, tc := range tt {
		tc.assertions(t, tc.args, tc.table)
	}
}
