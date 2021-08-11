package datetime

import (
	"testing"
	"time"

	"github.com/r2d2-ai/aiflow/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestFnSubHours_Eval(t *testing.T) {

	tests := []struct {
		Time     string
		Hours    int
		Expected string
	}{
		{
			Time:     "2020-03-19T15:02:03",
			Hours:    4,
			Expected: "2020-03-19T11:02:03Z",
		},
		{
			Time:     "2020-03-19T15:02:03",
			Hours:    3,
			Expected: "2020-03-19T12:02:03Z",
		},
	}

	in := &fnSubHours{}

	for _, d := range tests {
		final, err := in.Eval(d.Time, d.Hours)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final.(time.Time).Format(time.RFC3339))
	}
}
