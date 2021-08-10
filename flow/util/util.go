package util

import (
	"github.com/mohae/deepcopy"
	"github.com/r2d2-ai/AIflow/core/data/coerce"
)

func DeepCopy(data interface{}) interface{} {
	return deepcopy.Copy(data)
}

func DeepCopyMap(data map[string]interface{}) map[string]interface{} {
	copiedData := deepcopy.Copy(data)
	copiedMap, _ := coerce.ToObject(copiedData)
	return copiedMap
}
