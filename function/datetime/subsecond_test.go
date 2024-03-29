package datetime

import (
	"testing"
	"time"

	"github.com/r2d2-ai/aiflow/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func init() {
	function.ResolveAliases()
}

func TestFnSubSecond_Eval(t *testing.T) {

	tests := []struct {
		Time     string
		Seconds  int
		Expected string
	}{
		{
			Time:     "2020-03-19T15:02:03",
			Seconds:  10,
			Expected: "2020-03-19T15:01:53Z",
		},
		{
			Time:     "2020-03-19T15:02:03",
			Seconds:  30,
			Expected: "2020-03-19T15:01:33Z",
		},
	}

	in := &fnSubSecond{}

	for _, d := range tests {
		final, err := in.Eval(d.Time, d.Seconds)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final.(time.Time).Format(time.RFC3339))
	}
}
