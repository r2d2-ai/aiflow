package definition

import (
	"testing"

	"github.com/r2d2-ai/aiflow/data/expression"
	"github.com/r2d2-ai/aiflow/data/resolve"
	"github.com/stretchr/testify/assert"
)

func TestExprFactory(t *testing.T) {

	SetExprFactory(expression.NewFactory(resolve.GetBasicResolver()))

	assert.NotNil(t, GetExprFactory())

	linkExErr := NewLinkExprError("test")

	assert.NotNil(t, linkExErr)

	assert.Equal(t, "test", linkExErr.Error())
}

//func TestLinks(t *testing.T) {
//	defRep := &DefinitionRep{}
//
//	err := json.Unmarshal([]byte(defJSON), defRep)
//	assert.Nil(t, err)
//
//	def, err := NewDefinition(defRep)
//	assert.Nil(t, err)
//
//	assert.Nil(t, GetExpressionLinks(def))
//}
