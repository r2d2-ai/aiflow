package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnObjKeys(t *testing.T) {
	input := `{
		"key1": 123,
		"key2": {
		  "nested": 345.45,
		  "deepNesting": {
			"key3": 678.567575676,
			"key4": [
			  {
				"key5": 7987897.96878
			  }
			],
			"key5": [1, 2, 3]
		  }
		},
		"key3": [1, 2, 3],
		"key4": [
		  [1, 2],
		  [3, 4],
		  [3, 4, 5],
		  [
			3,
			4,
			{
			  "key8": 7987897.96878
			}
		  ]
		]
	  }`
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(input), &data)
	assert.Nil(t, err)
	f := &fnObjKeys{}
	v, err := f.Eval(data)
	assert.Nil(t, err)
	fmt.Println(v)
}
