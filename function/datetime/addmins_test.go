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

func TestFnAddMins_Eval(t *testing.T) {

	tests := []struct {
		Time     string
		Mins     int
		Expected string
	}{
		{
			Time:     "2020-03-19T15:02:03",
			Mins:     10,
			Expected: "2020-03-19T15:12:03Z",
		},
		{
			Time:     "2020-03-19T15:02:03",
			Mins:     30,
			Expected: "2020-03-19T15:32:03Z",
		},
	}

	in := &fnAddMins{}
	for _, d := range tests {
		final, err := in.Eval(d.Time, d.Mins)
		assert.Nil(t, err)
		assert.Equal(t, d.Expected, final.(time.Time).Format(time.RFC3339))
	}
}
